package app

import (
	"context"

	"github.com/hsndmr/go-sanctum/pkg/config"
	"github.com/hsndmr/go-sanctum/pkg/connection"
	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
	"github.com/hsndmr/go-sanctum/repositories"
)


type Container struct {
	Config *config.Config
	DBClient *connection.DBClient
	Repository *repositories.Repository
	Service *Service
}

type Service struct {
	Crypto *cryptoservice.Crypto
}

// C is the global container
var C *Container

// Init initializes the configuration and database connection
func Init() {
	C = &Container{}
	C.Config, _ = config.Init()
	C.DBClient, _ = connection.CreateClient(C.Config)
	C.Repository = repositories.CreateRepository(C.DBClient, context.Background())
	C.Service = CreateService()
}

func CreateService() *Service {
	return &Service{
		Crypto: &cryptoservice.Crypto{},
	}
}