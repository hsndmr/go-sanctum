package repositories

import (
	"context"
	"strings"
	"testing"
)

// TestCreatePersonalAccessToken tests the creation of a personal access token
func TestCreatePersonalAccessToken(t *testing.T) {
	uf := &UserFactory{}

	user := uf.Create()

	patr := &PersonalAccessTokenRepository{}

	t.Run("it should create personal access token", func(t *testing.T) {
		token, err :=  patr.Create(&CreatePersonalAccessTokenDto{
			User: user,
		}, context.Background())

		if err != nil {
			t.Errorf("failed creating personal access token: %s", err.Error())
		}

		if token.Token == "" {
			t.Errorf("failed creating personal access token")
		}
	})

	t.Run("it should create personal access token with abilities", func(t *testing.T) {
		abilities := []string{"user:update", "user:delete"}
		
		token, _ :=  patr.Create(&CreatePersonalAccessTokenDto{
			User: user,
			Abilities: abilities,
		}, context.Background())

		if *token.Abilities != strings.Join(abilities, ",") {
			t.Errorf("failed creating personal access token")
		}
	})
}