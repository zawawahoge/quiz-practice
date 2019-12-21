package adminserviceimpl

import (
	"context"
	"fmt"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

type adminServiceServer struct {
	service.AdminServiceServer
}

// New creats loginServiceServer implementation.
func New() service.AdminServiceServer {
	return &adminServiceServer{}
}

func (s *adminServiceServer) Hello(con context.Context, req *service.CreateAccountRequest) (*service.CreateAccountResponse, error) {
	id := req.GetId()
	password := req.GetPassword()
	fmt.Printf("id=%s, password=%s", id, password)
	return &service.CreateAccountResponse{
		Msg: "hello",
	}, nil
}
