package repositories

import (
	"context"

	"github.com/bxcodec/faker/v3"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/ent"
)

type UserFaker struct {
	Name string `faker:"name"`
	Email string `faker:"email"`
	Password string `faker:"password"`
}

type UserFactory struct {}

func (f *UserFactory) Create() (*ent.User, error) {
	item := UserFaker{}
	err := faker.FakeData(&item)
	
	if err != nil {
		return nil, err
	}

	ur := &UserRepository{}
	return ur.Create(&CreateUserDto{
		Name: item.Name,
		Email: item.Email,
		Password: item.Password,
	}, app.C.DBClient, context.Background())
}