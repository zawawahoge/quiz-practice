package adminserviceimpl

import (
	"context"
	"fmt"

	"github.com/zawawahoge/quiz-practice/api/app/domain/repository"
	"github.com/zawawahoge/quiz-practice/api/app/util/log"

	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
)

type adminServiceServer struct {
	service.AdminServiceServer
	accountRepository repository.AccountRepository
}

// New creats loginServiceServer implementation.
func New(accountRepository repository.AccountRepository) service.AdminServiceServer {
	return &adminServiceServer{
		accountRepository: accountRepository,
	}
}

func (s *adminServiceServer) CreateAccount(ctx context.Context, req *service.CreateAccountRequest) (*service.CreateAccountResponse, error) {
	logger := log.New(ctx)

	id := req.GetId()
	password := req.GetPassword()
	msg := fmt.Sprintf("id=%s, password=%s", id, password)
	logger.Debug("hello logger")
	return &service.CreateAccountResponse{
		Msg: msg,
	}, nil
}

func (s *adminServiceServer) GetAccounts(ctx context.Context, req *service.GetAccountsRequest) (*service.GetAccountsResponse, error) {
	logger := log.New(ctx)

	logger.Debug("get accounts")
	return &service.GetAccountsResponse{
		Accounts: []*service.GetAccountsResponse_Account{
			&service.GetAccountsResponse_Account{
				Id: "ID1",
			},
		},
	}, nil
}
