package main

import (
	"context"
	"log"
	"net"

	"github.com/danielsdev/microservices-go-gRPC/manager/pb"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "MANAGER, " + in.GetName()}, nil
}

func (s *Server) ListStudents(ctx context.Context, in *pb.ListStudentsRequest) (*pb.ListStudentsResponse, error) {
	return &pb.ListStudentsResponse{Name: "Test"}, nil
}

func main() {
	println("Running gRPC server")

	listener, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
