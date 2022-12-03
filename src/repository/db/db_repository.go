package db

import (
	"github.com/PaulTabaco/bookstore_oauth-api/src/domain/access_token"
	"github.com/PaulTabaco/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("db connection not implemented")
}
