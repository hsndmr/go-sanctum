package connection

import (
	"testing"

	"github.com/hsndmr/go-sanctum/pkg/config"
)

func TestCreateClient(t *testing.T) {
	config := &config.Config{
		Database: &config.Database{
			Connection: "sqlite3",
			Host:       "localhost",
			Name:       ":memory:",
			User:       "root",
			Password:   "",
		},
	}

	_, err:= CreateClient(config)

	if err !=nil {
		t.Errorf("CreateClient error: %v", err)
	}
}