package main

import (
	"log"
	"net"

	"github.com/zawawahoge/quiz-practice-api/app/infrastructure/impl/loginserviceimpl"
	"github.com/zawawahoge/quiz-practice-api/app/proto/v1/loginservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	loginservice.RegisterLoginServiceServer(s, loginserviceimpl.New())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
