package repository

import (
	"github.com/zawawahoge/quiz-practice/api/app/domain/model"
)

// AccountRepository is an interface of account repository.
type AccountRepository interface {
	Create(model.AccountID, model.AccountPasswordHash) (*model.Account, error)
	List() ([]*model.Account, error)
}
