package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mubashir7933/RestAPI-Golang/internal/config"
)

func main() {
	//load config file
	cfg := config.MustLoad()

	// database setup
	// setup router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students API"))
	})
	// setup Server

	server := http.Server{
		Addr:    cfg.Adress,
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server: ", err.Error())
	}
	fmt.Println("Server started on port", cfg.Adress)
}
