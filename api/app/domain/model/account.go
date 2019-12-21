package model

// AccountID is a type of account id.
type AccountID string

// AccountPasswordHash is a hashed password of account.
type AccountPasswordHash string

// Account is a model of account.
type Account struct {
	ID           AccountID `gorm:"primary_key"`
	PasswordHash AccountPasswordHash
}
