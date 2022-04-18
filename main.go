package main

import (
	"net/http"

	"github.com/hsndmr/go-sanctum/app"
	"github.com/hsndmr/go-sanctum/routers"
)


func main() {

	app.Init()
	defer app.C.DBClient.Close()

	handler := routers.InitRouter()

	server := &http.Server{
		Addr:           ":3000",
		Handler:        handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}