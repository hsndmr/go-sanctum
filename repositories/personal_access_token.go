package repositories

import (
	"context"
	"time"

	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/connection"
	"github.com/hsndmr/go-sanctum/pkg/token"
)

type CreatePersonalAccessTokenDto struct {
	Name string
	User *ent.User
	ExpirationAt *time.Time
	Abilities []string
}

type PersonalAccessTokenRepository struct {
	*BaseRepository
	Token *token.Token
}
// CreatePersonalAccessToken creates a new personal access token
func (p *PersonalAccessTokenRepository) Create(dto *CreatePersonalAccessTokenDto, db *connection.DBClient, ctx context.Context) (*ent.PersonalAccessToken, error) {
	expirationAt := time.Now().Add(time.Hour * 24 * 7)
	if(dto.ExpirationAt != nil) {
		expirationAt = *dto.ExpirationAt
	}

	tokenItem, err := p.Token.Create()

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