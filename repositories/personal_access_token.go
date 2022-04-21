package repositories

import (
	"time"

	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/token"
)

type CreatePersonalAccessTokenDto struct {
	Name string
	User *ent.User
	ExpirationAt *time.Time
	Abilities []string
}

type PersonalAccessTokenRepositoryI interface {
	Create(dto *CreatePersonalAccessTokenDto) (*ent.PersonalAccessToken, error)
}
type PersonalAccessTokenRepository struct {
	*BaseRepository
	Token token.TokenI
}
// CreatePersonalAccessToken creates a new personal access token
func (p *PersonalAccessTokenRepository) Create(dto *CreatePersonalAccessTokenDto) (*ent.PersonalAccessToken, error) {
	expirationAt := time.Now().Add(time.Hour * 24 * 7)
	if(dto.ExpirationAt != nil) {
		expirationAt = *dto.ExpirationAt
	}

	tokenItem, err := p.Token.Create()

	if err != nil {
		return nil, err
	}
	
	personalAccessTokenCreate := p.db.Client().PersonalAccessToken.Create()

	return personalAccessTokenCreate.
					SetName(dto.Name).
					SetUser(dto.User).
					SetToken(tokenItem.Hash).
					SetExpirationAt(expirationAt).
					SetAbilities(dto.Abilities).
					Save(p.ctx)
}