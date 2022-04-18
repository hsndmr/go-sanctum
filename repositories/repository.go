package repositories

import (
	"context"

	"github.com/hsndmr/go-sanctum/pkg/connection"
)

type BaseRepository struct {
	db  *connection.DBClient
	ctx context.Context
}


type Repository struct {
	User                *UserRepository
	PersonalAccessToken *PersonalAccessTokenRepository
}


func CreateRepository(db *connection.DBClient, ctx context.Context) *Repository {

	bs := &BaseRepository{
		db: db,
		ctx: ctx,
	}

	return &Repository{
		User: &UserRepository{bs},
		PersonalAccessToken: &PersonalAccessTokenRepository{bs},
	}
}
