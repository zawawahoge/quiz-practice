package loginserviceimpl

import (
	"context"

	"github.com/zawawahoge/quiz-practice/api/app/domain/repository"
	"github.com/zawawahoge/quiz-practice/api/app/util/log"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

type loginServiceServer struct {
	service.LoginServiceServer
	accountRepository repository.AccountRepository
}

// New creats loginServiceServer implementation.
func New(accountRepository repository.AccountRepository) service.LoginServiceServer {
	return &loginServiceServer{
		accountRepository: accountRepository,
	}
}

func (s *loginServiceServer) Hello(ctx context.Context, req *service.HelloRequest) (*service.HelloResponse, error) {
	logger := log.New(ctx)

	logger.Debug("hello logger")
	return &service.HelloResponse{
		Msg: "helloa",
	}, nil
}
