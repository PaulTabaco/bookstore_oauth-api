package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "new token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new token should not have defined access id")
	assert.True(t, at.UserId == 0, "new token should not have an assosiated userId")
	assert.EqualValues(t, 0, at.ClientId, "new token should not have an assosiated clientId")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token created three hours from now should not be expired by default")
}
