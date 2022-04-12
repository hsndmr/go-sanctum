package connection

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hsndmr/go-sanctum/ent"
	"github.com/hsndmr/go-sanctum/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

// GetDb returns the database connection
var Client *ent.Client

// initDb initializes the database connection
func initDb() {
	var err error
	Client, err = ent.Open(config.App.Database.Connection, createConnectionString())
	
	if err != nil {
		log.Fatalf("Database Connection Error: %v", err)
	}

}

func createConnectionString() string { 
	if(config.App.Database.Connection == "sqlite3") {
		return "file:ent?mode=memory&cache=shared&_fk=1"
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.App.Database.User,
		config.App.Database.Password,
		config.App.Database.Host,
		config.App.Database.Name,
	)
}

// connect connects to the all connections
func Connect() {
	initDb()
}

func Close() error {
	return Client.Close()
}