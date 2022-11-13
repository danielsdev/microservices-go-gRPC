package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/danielsdev/microservices-go-gRPC/api/pb"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	conn   *grpc.ClientConn
	err    error
	client pb.HelloClient
)

func initGrpcConnection() {
	if conn == nil {
		conn, err := grpc.Dial("localhost:5555", grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatal(err)
		}

		client = pb.NewHelloClient(conn)
	}
}

func ExibeTodosAlunos(c *gin.Context) {
	initGrpcConnection()
	req := &pb.ListStudentsRequest{}

	res, err := client.ListStudents(context.Background(), req)

	if err != nil {
		c.JSON(400, err)
	}

	c.JSON(200, res.GetStudent())
}

func CriaNovoAluno(c *gin.Context) {
	initGrpcConnection()

	type Request struct {
		Nome      string
		Matricula string
		Rg        string
		Cpf       string
		Id        string
	}

	var requestBody Request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	createAluno := &pb.StudentRequest{
		Name:      requestBody.Nome,
		Matricula: requestBody.Matricula,
		Rg:        requestBody.Rg,
		Cpf:       requestBody.Cpf,
	}

	res, err := client.CreateStudent(context.Background(), createAluno)
	if err != nil || res.Id == "0" {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusCreated, res)
}

func BuscaAlunoPorID(c *gin.Context) {

	initGrpcConnection()

	id := c.Params.ByName("id")

	req := &pb.GetStudentRequest{
		Id: id,
	}

	res, err := client.GetStudent(context.Background(), req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if res.Id == "0" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeletaAluno(c *gin.Context) {
	initGrpcConnection()

	id := c.Params.ByName("id")

	req := &pb.DeleteStudentRequest{
		Id: id,
	}

	res, err := client.DeleteStudent(context.Background(), req)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if res.Status == "Erro" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func EditaAluno(c *gin.Context) {
	initGrpcConnection()

	type Request struct {
		Nome      string
		Matricula string
		Rg        string
		Cpf       string
		Id        string
	}

	var requestBody Request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	editAluno := &pb.EditStudentRequest{
		Id:        c.Params.ByName("id"),
		Name:      requestBody.Nome,
		Matricula: requestBody.Matricula,
		Rg:        requestBody.Rg,
		Cpf:       requestBody.Cpf,
	}

	res, err := client.EditStudent(context.Background(), editAluno)
	if err != nil || res.Id == "0" {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
