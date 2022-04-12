package connection

import (
	"testing"

	"github.com/hsndmr/go-sanctum/pkg/config"
)

func TestInit(t *testing.T) {
	config.App.Init()
	Connect()
	defer Client.Close()
}