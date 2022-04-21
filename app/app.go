package app

import (
	"context"

	"github.com/hsndmr/go-sanctum/pkg/config"
	"github.com/hsndmr/go-sanctum/pkg/connection"
	cryptoservice "github.com/hsndmr/go-sanctum/pkg/crypto"
	"github.com/hsndmr/go-sanctum/pkg/services"
	"github.com/hsndmr/go-sanctum/pkg/token"
	"github.com/hsndmr/go-sanctum/repositories"
)


type Container struct {
	Config *config.Config
	DBClient connection.DBClientI
	Repository *repositories.Repository
	Service *services.Service
}



// C is the global container
var C *Container

// Init initializes the configuration and database connection
func Init() {
	C = &Container{}
	C.Config, _ = config.Init()
	C.DBClient, _ = connection.CreateClient(C.Config)
	C.Service = CreateService()
	C.Repository = repositories.CreateRepository(C.Service, C.DBClient, context.Background())
}

func CreateService() *services.Service {

	crypto := &cryptoservice.Crypto{}

	return &services.Service{
		Crypto: crypto,
		Token: &token.Token{
			Crypto: crypto,
		},
	}
}