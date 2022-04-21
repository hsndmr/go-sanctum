package repositories

import (
	"context"

	"github.com/hsndmr/go-sanctum/pkg/connection"
	"github.com/hsndmr/go-sanctum/pkg/services"
)

type BaseRepository struct {
	db  connection.DBClientI
	ctx context.Context
}

type Repository struct {
	User                UserRepositoryI
	PersonalAccessToken PersonalAccessTokenRepositoryI
}


func CreateRepository(services *services.Service, db connection.DBClientI, ctx context.Context) *Repository {

	bs := &BaseRepository{
		db: db,
		ctx: ctx,
	}

	return &Repository{
		User: &UserRepository{
			BaseRepository: bs,
			Crypto: services.Crypto,
		},
		PersonalAccessToken: &PersonalAccessTokenRepository{
			BaseRepository: bs,
			Token: services.Token,
		},
	}
}
