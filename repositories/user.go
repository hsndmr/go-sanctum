package repositories

import (
	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/ent/user"
	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
)

// dtos
type CreateUserDto struct {
	Name string
	Email string
	Password string
}

type UserRepositoryI interface {
	Create(dto *CreateUserDto) (*ent.User, error)
	FindByEmail(email string) (*ent.User, error)
}

type UserRepository struct {
	*BaseRepository
	Crypto *cryptoservice.Crypto
}

// Create creates a user and returns the user.
func (r *UserRepository) Create(dto *CreateUserDto) (*ent.User, error) {
	return r.db.Client().User.
						Create().
						SetEmail(dto.Email).
						SetPassword(r.Crypto.HashPassword(dto.Password)).
						SetName(dto.Name).
						Save(r.ctx)
}

// FindByEmail finds a user by email and returns the user.
func (r *UserRepository) FindByEmail(email string) (*ent.User, error) {
	return r.db.Client().User.
						Query().
						Where(user.Email(email)).
						Only(r.ctx)
}