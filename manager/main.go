package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/danielsdev/microservices-go-gRPC/manager/database"
	"github.com/danielsdev/microservices-go-gRPC/manager/models"
	"github.com/danielsdev/microservices-go-gRPC/manager/pb"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "MANAGER, " + in.GetName()}, nil
}

func (s *Server) CreateStudent(ctx context.Context, in *pb.StudentRequest) (*pb.StudentResponse, error) {
	aluno := models.Aluno{
		Nome:      in.GetName(),
		Matricula: in.GetMatricula(),
		CPF:       in.GetCpf(),
		RG:        in.GetRg(),
	}

	database.DB.Create(&aluno)

	id := strconv.FormatUint(uint64(aluno.ID), 10)

	return &pb.StudentResponse{Id: id, Name: aluno.Nome, Cpf: aluno.CPF, Rg: aluno.RG, Matricula: aluno.Matricula}, nil
}

func (s *Server) EditStudent(ctx context.Context, in *pb.EditStudentRequest) (*pb.StudentResponse, error) {
	var aluno models.Aluno

	id := in.GetId()

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		return &pb.StudentResponse{Id: "0"}, nil
	}

	if in.GetName() != "" {
		aluno.Nome = in.GetName()
	}

	if in.GetMatricula() != "" {
		aluno.Matricula = in.GetMatricula()
	}

	if in.GetCpf() != "" {
		aluno.CPF = in.GetCpf()
	}

	if in.GetRg() != "" {
		aluno.RG = in.GetRg()
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	returnId := strconv.FormatUint(uint64(aluno.ID), 10)

	return &pb.StudentResponse{Id: returnId, Name: aluno.Nome, Cpf: aluno.CPF, Rg: aluno.RG, Matricula: aluno.Matricula}, nil
}

func (s *Server) ListStudents(ctx context.Context, in *pb.ListStudentsRequest) (*pb.ListStudentsResponse, error) {
	var alunos []models.Aluno

	var alunos_response []*pb.StudentResponse

	database.DB.Find(&alunos)

	for _, element := range alunos {
		alunos_response = append(alunos_response, &pb.StudentResponse{
			Id:        strconv.FormatUint(uint64(element.ID), 10),
			Name:      element.Nome,
			Cpf:       element.CPF,
			Rg:        element.RG,
			Matricula: element.Matricula,
		})
	}

	return &pb.ListStudentsResponse{Student: alunos_response}, nil
}

func (s *Server) GetStudent(ctx context.Context, in *pb.GetStudentRequest) (*pb.StudentResponse, error) {
	var aluno models.Aluno
	requestId := in.GetId()

	database.DB.First(&aluno, requestId)

	id := strconv.FormatUint(uint64(aluno.ID), 10)

	return &pb.StudentResponse{Id: id, Name: aluno.Nome, Cpf: aluno.CPF, Rg: aluno.RG, Matricula: aluno.Matricula}, nil
}

func (s *Server) DeleteStudent(ctx context.Context, in *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	var aluno models.Aluno
	requestId := in.GetId()

	err := database.DB.Delete(&aluno, requestId).Error

	status := "Ok"

	if err != nil || requestId == "" {
		status = "Erro"
	}

	return &pb.DeleteStudentResponse{Status: status}, nil
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panic("Erro ao carregar env")
	}

	database.ConectaComBancoDeDados()

	println("Running gRPC server")

	listener, err := net.Listen("tcp", "localhost:"+os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
