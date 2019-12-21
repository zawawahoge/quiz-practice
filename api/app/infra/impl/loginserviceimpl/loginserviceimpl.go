package loginserviceimpl

import (
	"context"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

type loginServiceServer struct {
	service.LoginServiceServer
}

// New creats loginServiceServer implementation.
func New() service.LoginServiceServer {
	return &loginServiceServer{}
}

func (s *loginServiceServer) Hello(con context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	return &service.HelloResponse{
		Msg: "hello",
	}, nil
}
