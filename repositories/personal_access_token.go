package repositories

import (
	"context"
	"time"

	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/connection"
	"github.com/hsndmr/go-sanctum/pkg/token"
)

type PersonalAccessTokenI interface{
	Create(dto *CreatePersonalAccessTokenDto, db *connection.DBClient, ctx context.Context) (*ent.PersonalAccessToken, error)
} 

type CreatePersonalAccessTokenDto struct {
	Name string
	User *ent.User
	ExpirationAt *time.Time
	Abilities []string
}

type PersonalAccessTokenRepository struct {
	PersonalAccessTokenI
}
// CreatePersonalAccessToken creates a new personal access token
func (p *PersonalAccessTokenRepository) Create(dto *CreatePersonalAccessTokenDto, db *connection.DBClient, ctx context.Context) (*ent.PersonalAccessToken, error) {

	expirationAt := time.Now().Add(time.Hour * 24 * 7)
	if(dto.ExpirationAt != nil) {
		expirationAt = *dto.ExpirationAt
	}

	tokenItem, err := token.CreateToken()

	if err != nil {
		return nil, err
	}
	
	personalAccessTokenCreate := db.Client.PersonalAccessToken.Create()

	return personalAccessTokenCreate.
					SetName(dto.Name).
					SetUser(dto.User).
					SetToken(tokenItem.Hash).
					SetExpirationAt(expirationAt).
					SetAbilities(dto.Abilities).
					Save(ctx)
}