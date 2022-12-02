package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("new token should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("new token should not have defined access id")
	}

	if at.UserId != 0 {
		t.Error("new token should not have an assosiated userId")
	}

	if at.ClientId != 0 {
		t.Error("new token should not have an assosiated clientId")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token created three hours from now should not be expired by default")
	}

}
