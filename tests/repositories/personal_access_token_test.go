package repositorytest

import (
	"testing"
	"time"

	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/repositories"
	"github.com/stretchr/testify/assert"
)

// TestCreatePersonalAccessToken tests the creation of a personal access token
func TestCreatePersonalAccessToken(t *testing.T) {
	uf := &UserFactory{}

	user, _ := uf.Create()

	patr := app.C.Repository.PersonalAccessToken

	t.Run("it should create personal access token", func(t *testing.T) {
		token, _, err :=  patr.Create(&repositories.CreatePersonalAccessTokenDto{
			User: user,
		})

		if err != nil {
			t.Errorf("failed creating personal access token: %s", err.Error())
		}

		if token.Token == "" {
			t.Errorf("failed creating personal access token")
		}
	})

	t.Run("it should create personal access token with abilities", func(t *testing.T) {
		abilities := []string{"user:update", "user:delete"}
		
		 token, _, _ := patr.Create(&repositories.CreatePersonalAccessTokenDto{
			User: user,
			Abilities: abilities,
		})
		
		assert.Equal(t, token.Abilities, abilities, "the abilities should be equal")
	})
}

func TestVerifyToken(t *testing.T) {
	uf := &UserFactory{}

	user, _ := uf.Create()

	patr := app.C.Repository.PersonalAccessToken

	t.Run("it should verify token", func(t *testing.T) {
		_, plainText, _ := patr.Create(&repositories.CreatePersonalAccessTokenDto{
			User: user,
		})

		_, err := patr.Verify(plainText)

		if err != nil {
			t.Errorf("failed verifying token: %s", err.Error())
		}
	})

	t.Run("it should throw expired token error", func(t *testing.T) {
		time := time.Now().Add(-time.Hour*1)
		_, plainText, _ := patr.Create(&repositories.CreatePersonalAccessTokenDto{
			User: user,
			ExpirationAt: &time,
		})

		_, err := patr.Verify(plainText)

		if err == nil {
			t.Errorf("not throw expired token error")
		}
	})
}