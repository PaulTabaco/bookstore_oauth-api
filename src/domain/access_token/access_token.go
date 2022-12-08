package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/PaulTabaco/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/PaulTabaco/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"` // Web, Android, iOs etc.
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expires id")
	}
	return nil
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	// switch at.GrantType {
	// case grantTypePassword:
	// 	break

	// case grandTypeClientCredentials:
	// 	break

	// default:
	// 	return errors.NewBadRequestError("invalid grant_type parameter")
	// }

	//TODO: Validate parameters for each grant_type
	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
