package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/routers"
)


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