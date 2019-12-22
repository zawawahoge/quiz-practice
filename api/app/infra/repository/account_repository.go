package database

import (
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
func (repo *dbAccountRepository) Create(accountID model.AccountID, accountPasswordHash model.AccountPasswordHash) (*model.Account, error) {
	account := &model.Account{
		ID:           accountID,
		PasswordHash: accountPasswordHash,
	}
	err := repo.db.Create(account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (repo *dbAccountRepository) List() ([]*model.Account, error) {
	var accounts []*model.Account
	err := repo.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
