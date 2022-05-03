package repositorytest

import (
	"testing"

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