package app

import (
	"github.com/hsndmr/go-sanctum/pkg/config"
	"github.com/hsndmr/go-sanctum/pkg/connection"
)


type Container struct {
	Config *config.Config
	DBClient *connection.DBClient
}

// C is the global container
var C *Container

// Init initializes the configuration and database connection
func Init() {
	C := &Container{}
	C.Config, _ = config.Init()
	C.DBClient, _ = connection.CreateClient(C.Config)
}