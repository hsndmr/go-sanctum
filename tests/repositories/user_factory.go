package repositorytest

import (
	"github.com/bxcodec/faker/v3"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/repositories"
)

type UserFaker struct {
	Name string `faker:"name"`
	Email string `faker:"email"`
}

type UserFactory struct {}

func (f *UserFactory) Create() (*ent.User, error) {
	item := UserFaker{}
	err := faker.FakeData(&item)
	
	if err != nil {
		return nil, err
	}

	ur := app.C.Repository.User

	
	return ur.Create(&repositories.CreateUserDto{
		Name: item.Name,
		Email: item.Email,
		Password: "secret",
	})
}