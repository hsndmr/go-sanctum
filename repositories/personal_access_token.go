package repositories

import (
	"errors"
	"strconv"
	"time"

	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/ent/personalaccesstoken"
	"github.com/hsndmr/go-sanctum/pkg/token"
)

type CreatePersonalAccessTokenDto struct {
	Name string
	User *ent.User
	ExpirationAt *time.Time
	Abilities []string
}

type PersonalAccessTokenRepositoryI interface {
	Create(dto *CreatePersonalAccessTokenDto) (*ent.PersonalAccessToken, string, error)
	Verify(token string) (*ent.User, error)
}
type PersonalAccessTokenRepository struct {
	*BaseRepository
	Token token.TokenI
	UserRepository UserRepositoryI
}
// CreatePersonalAccessToken creates a new personal access token
func (p *PersonalAccessTokenRepository) Create(dto *CreatePersonalAccessTokenDto) (*ent.PersonalAccessToken, string, error) {
	expirationAt := time.Now().Add(time.Hour * 24 * 7)
	if(dto.ExpirationAt != nil) {
		expirationAt = *dto.ExpirationAt
	}

	tokenItem, err := p.Token.Create()

	if err != nil {
		return nil, "", err
	}
	
	personalAccessTokenCreate := p.db.Client().PersonalAccessToken.Create()

	personalAccessToken, err:= personalAccessTokenCreate.
		SetName(dto.Name).
		SetUser(dto.User).
		SetToken(tokenItem.Hash).
		SetExpirationAt(expirationAt).
		SetAbilities(dto.Abilities).
		Save(p.ctx)
	
	if err != nil {
		return nil, "", err
	}

	return personalAccessToken, tokenItem.GetPlainText(strconv.Itoa(personalAccessToken.ID)), nil
}

func (p *PersonalAccessTokenRepository) Verify(token string) (*ent.User, error) {
	idStr, hash, err := p.Token.Split(token)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, err
	}
	
	personalAccessToken, err := p.db.Client().PersonalAccessToken.
		Query().
		Where(personalaccesstoken.ID(id)).
		Where(personalaccesstoken.Token(hash)).
		Only(p.ctx)

	if err != nil {
		return nil, err
	}

	if personalAccessToken.ExpirationAt.Before(time.Now()) {
		return nil, errors.New("token expired") 
	}

	user, err := p.UserRepository.FindByID(personalAccessToken.UserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}