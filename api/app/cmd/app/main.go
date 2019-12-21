package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/zawawahoge/quiz-practice/api/app/infrastructure/impl/adminserviceimpl"
	"github.com/zawawahoge/quiz-practice/api/app/infrastructure/impl/loginserviceimpl"
	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

const (
	port = ":9090"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	service.RegisterLoginServiceServer(s, loginserviceimpl.New())
	service.RegisterAdminServiceServer(s, adminserviceimpl.New())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
