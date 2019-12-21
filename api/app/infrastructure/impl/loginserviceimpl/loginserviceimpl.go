package loginserviceimpl

import (
	"context"
	"fmt"
	"log"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/loginservice"
)

type loginServiceServer struct {
	loginservice.LoginServiceServer
}

// New creats loginServiceServer implementation.
func New() loginservice.LoginServiceServer {
	return &loginServiceServer{}
}

func (s *loginServiceServer) Login(con context.Context, req *loginservice.LoginRequest) (*loginservice.LoginResponse, error) {
	email := req.GetEmail()
	password := req.GetPassword()
	return &loginservice.LoginResponse{
		Token: email + password,
	}, nil
}

func (s *loginServiceServer) Hello(con context.Context, req *loginservice.HelloRequest) (*loginservice.HelloResponse, error) {
	var Logger *log.Logger
	Logger.Println("Hello")
	fmt.Println("hello fmt")
	return &loginservice.HelloResponse{
		Msg: "hello",
	}, nil
}
