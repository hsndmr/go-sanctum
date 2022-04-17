package connection

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

type DBClient struct {
	Client *ent.Client
}

// CreateScheme creates the scheme
func (c *DBClient) CreateScheme(ctx context.Context) error {
	return c.Client.Schema.Create(ctx)
}

// Close closes the client
func (c *DBClient) Close() error {
	return c.Client.Close()
}

// CreateClient creates the client to connect db
func CreateClient(config *config.Config) (*DBClient, error) {
	client, err := ent.Open(config.Database.Connection, createConnectionString(config))
	
	if err != nil {
		return nil, err
	}

	dbClient := &DBClient{
		Client: client,
	}

	if (config.EnvType == "local") {
		dbClient.CreateScheme(context.Background())
	}

	return dbClient, nil
}

// createConnectionString creates the connection string
func createConnectionString(config *config.Config) string { 
	if(config.Database.Connection == "sqlite3") {
		return "file:ent?mode=memory&cache=shared&_fk=1"
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Name,
	)
}