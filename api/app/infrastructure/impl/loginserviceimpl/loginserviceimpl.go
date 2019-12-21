package loginserviceimpl

import (
	"context"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

type loginServiceServer struct {
	loginservice.LoginServiceServer
}

// New creats loginServiceServer implementation.
func New() loginservice.LoginServiceServer {
	return &loginServiceServer{}
}

func (s *loginServiceServer) Hello(con context.Context, req *loginservice.HelloRequest) (*loginservice.HelloResponse, error) {
	return &loginservice.HelloResponse{
		Msg: "hello",
	}, nil
}
