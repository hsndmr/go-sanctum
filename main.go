package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/routers"
)

// @title           Go Sanctum
// @version         1.0
// @description     Go Sanctum offers a simple-to-use authentication system like Laravel Sanctum.
// @host      localhost:3000
// @BasePath  /api/v1
func main() {

	app.Init()
	defer app.C.DBClient.Close()

	if !app.C.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := routers.InitRouter()

	server := &http.Server{
		Addr:           ":"+app.C.Config.Port,
		Handler:        handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}