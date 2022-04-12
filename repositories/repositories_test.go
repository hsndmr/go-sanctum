package repositories

import (
	"context"
	"os"
	"testing"

	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/pkg/connection"
)


func TestMain(m *testing.M) {
	app.Init(context.Background())
	exitVal := m.Run()
	connection.Close()
	os.Exit(exitVal)
}