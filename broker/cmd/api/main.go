package main

import (
	"log"
	"net/http"
)

const (
	PORT = "80"
)

type Config struct {}

func main() {
	app := Config{}
	hanler := app.routes()

	log.Println("Starting broker service on port " + PORT)

	srv := &http.Server{
		Addr: ":" + PORT,
		Handler: hanler,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}