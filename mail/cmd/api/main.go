package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	PORT = "80"
)

type Config struct {
	Mailer Mail
}

func main() {
	app := Config{
		Mailer: createMail(),
	}

	log.Println("Starting mail service on port ", PORT)

	srv := &http.Server{
		Addr: ":" + PORT,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)	
	}
}

func createMail() Mail {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Panic(err)
	}

	m := Mail{
		Domain: os.Getenv("MAIL_DOMAIN"),
		Host: os.Getenv("MAIL_HOST"),
		Port: port,
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
		Encryption: os.Getenv("MAIL_ENCRYPTION"),
		FromName: os.Getenv("MAIL_FROMNAME"),
		FromAddress: os.Getenv("MAIL_FROMADDRESS"),
	}

	return m
}