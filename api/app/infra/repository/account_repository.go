package database

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/zawawahoge/quiz-practice/api/app/domain/model"
	"github.com/zawawahoge/quiz-practice/api/app/domain/repository"
)

type dbAccountRepository struct {
	db *gorm.DB
	repository.AccountRepository
}

// NewAccountRepository creates a new repository for account.
func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &dbAccountRepository{
		db: db,
	}
}

// Create saves an account to DB.
func (repo *dbAccountRepository) Create(ctx context.Context, accountID model.AccountID, accountPasswordHash model.AccountPasswordHash) (*model.Account, error) {
	account := &model.Account{
		ID:           accountID,
		PasswordHash: accountPasswordHash,
	}
	repo.db.Create(account)
	return account, nil
}
