package loginserviceimpl

import (
	"context"
	"github.com/zawawahoge/quiz-practice/api/app/util/log"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

type loginServiceServer struct {
	service.LoginServiceServer
}

// New creats loginServiceServer implementation.
func New() service.LoginServiceServer {
	return &loginServiceServer{}
}

func (s *loginServiceServer) Hello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	logger := log.New(ctx)

	logger.Debug("hello logger")
	return &service.HelloResponse{
		Msg: "helloa",
	}, nil
}
