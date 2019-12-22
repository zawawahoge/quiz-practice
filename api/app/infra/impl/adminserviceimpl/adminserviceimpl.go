package adminserviceimpl

import (
	"context"
	"fmt"

	"github.com/zawawahoge/quiz-practice/api/app/domain/model"
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

	id := model.AccountID(req.GetId())
	password := model.AccountPasswordHash(req.GetPassword())
	msg := fmt.Sprintf("id=%s, password=%s", id, password)
	logger.Debug("Debug")
	account, err := s.accountRepository.Create(id, password)
	if err != nil {
		return nil, err
	}
	logger.Debug(fmt.Sprintf("CreateAccount successfully; account=%#v", account))
	return &service.CreateAccountResponse{
		Msg: msg,
	}, nil
}

func (s *adminServiceServer) GetAccounts(ctx context.Context, req *service.GetAccountsRequest) (*service.GetAccountsResponse, error) {
	logger := log.New(ctx)

	logger.Debug("get accounts")
	accounts, err := s.accountRepository.List()
	if err != nil {
		logger.Debug(fmt.Sprintf("failed to get accounts; err=%v", err))
		return nil, err
	}
	accountIds := make([]*service.GetAccountsResponse_Account, 0, len(accounts))
	for _, v := range accounts {
		accountIds = append(accountIds, &service.GetAccountsResponse_Account{
			Id: string(v.ID),
		})
	}
	return &service.GetAccountsResponse{Accounts: accountIds}, nil
}
