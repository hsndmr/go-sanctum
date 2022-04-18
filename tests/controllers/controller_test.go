package controllertest

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/routers"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	app.Init()
	defer app.C.DBClient.Client.Close()
	router = routers.InitRouter()
	exitVal := m.Run()
	os.Exit(exitVal)
}

// jsonReaderFactory creates a json reader from struct
func jsonReaderFactory(in interface{}) (io.Reader) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.Encode(in)

	return buf
}