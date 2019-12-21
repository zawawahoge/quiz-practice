package adminserviceimpl

import (
	"github.com/zawawahoge/quiz-practice/api/app/domain/model"
	"github.com/zawawahoge/quiz-practice/api/app/proto/v1/service"
	"testing"
)

func Test_AdminService_CreateAccount(t *testing.T) {
	req := &service.CreateAccountRequest{
		Id:       model.AccountID("TEST_ID"),
		Password: model.AccountPasswordHash("TEST_PASSWORD"),
	}
	gomock

}
