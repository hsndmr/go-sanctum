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
	FindByID(id int) (*ent.User, error)
	ToInterfaceForJson(u *ent.User) interface{}
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

// FindById finds a user by id and returns the user.
func (r *UserRepository) FindByID(id int) (*ent.User, error) {
	return r.db.Client().User.
		Query().
		Where(user.ID(id)).
		Only(r.ctx)
}

func (r *UserRepository) ToInterfaceForJson(u *ent.User) interface{} {
	return struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Email string `json:"email"`
	} {
		Id: u.ID,
		Name: u.Name,
		Email: u.Email,
	}
}