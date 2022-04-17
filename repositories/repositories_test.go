package repositories

import (
	"os"
	"testing"

	"github.com/hsndmr/go-sanctum/app"
)


func TestMain(m *testing.M) {
	app.Init()
	exitVal := m.Run()
	app.C.DBClient.Client.Close()
	os.Exit(exitVal)
}