package app

import (
	"context"
	"log"

	"github.com/hsndmr/go-sanctum/pkg/config"
	"github.com/hsndmr/go-sanctum/pkg/connection"
)

func Init(ctx context.Context) {
	config.App.Init()
	connection.Connect()

	if config.App.Env == "local" {
		CreateSchema()
	}
}

func CreateSchema() {
	if err := connection.Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Database Migration Error: %v", err)
	}
}