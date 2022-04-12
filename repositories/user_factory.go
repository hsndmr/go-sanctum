package repositories

import (
	"context"

	"github.com/bxcodec/faker/v3"
	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/helpers"
)

type UserFaker struct {
	Name string `faker:"name"`
	Email string `faker:"email"`
	Password string `faker:"password"`
}

type UserFactory struct {}

func (f *UserFactory) Create() (*ent.User) {
	item := UserFaker{}
	err := faker.FakeData(&item)
	helpers.CheckErr(err)
	
	ur := &UserRepository{}
	u, err := ur.Create(&CreateUserDto{
		Name: item.Name,
		Email: item.Email,
		Password: item.Password,
	}, context.Background())
	helpers.CheckErr(err)

	return u
}