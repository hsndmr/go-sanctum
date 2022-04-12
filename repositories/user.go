package repositories

import (
	"context"

	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/connection"
	"github.com/hsndmr/go-sanctum/pkg/crypto"
)

type CreateUserDto struct {
	Name string
	Email string
	Password string
}

type UserRepository struct {

}

// Create creates a user and returns the user.
func (r *UserRepository) Create(dto *CreateUserDto, ctx context.Context) (*ent.User, error) {
	return connection.Client.User.
						Create().
						SetEmail(dto.Email).
						SetPassword(crypto.HashPassword(dto.Password)).
						SetName(dto.Name).
						Save(ctx)
}