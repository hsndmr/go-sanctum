package app

import (
	"context"

	"github.com/hsndmr/go-sanctum/pkg/config"
	"github.com/hsndmr/go-sanctum/pkg/connection"
)


type Container struct {
	Config *config.Config
	DBClient *connection.DBClient
}

var C *Container

// Init initializes the configuration and database connection
func Init() {
	C = CreateContainer(context.Background())
}

func CreateContainer(ctx context.Context) *Container {
	container := &Container{}
	container.Config, _ = config.Init()
	container.DBClient, _ = connection.CreateClient(container.Config)
	return container
}