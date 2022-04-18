package repositorytest

import (
	"os"
	"testing"

	"github.com/hsndmr/go-sanctum/app"
)


func TestMain(m *testing.M) {
	app.Init()
	defer app.C.DBClient.Client.Close()
	exitVal := m.Run()
	os.Exit(exitVal)
}