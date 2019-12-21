package repository

import (
	"context"

	"github.com/zawawahoge/quiz-practice/api/app/domain/model"
)

// AccountRepository is an interface of account repository.
type AccountRepository interface {
	Create(context.Context, model.AccountID, model.AccountPasswordHash) (*model.Account, error)
}
