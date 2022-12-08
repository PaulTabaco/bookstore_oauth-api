package access_token

import (
	"strings"

	"github.com/PaulTabaco/bookstore_oauth-api/src/domain/access_token"
	"github.com/PaulTabaco/bookstore_oauth-api/src/repository/db"
	"github.com/PaulTabaco/bookstore_oauth-api/src/repository/rest"
	"github.com/PaulTabaco/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.RestUsersRepository
	dbRepo        db.DbRepository
}

func NewService(userRepo rest.RestUsersRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: userRepo,
		dbRepo:        dbRepo,
	}
}

func (s service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	// User authentification by Users API
	user, err := s.restUsersRepo.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	// New access token generation
	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	// Save this new token in cassandra
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
