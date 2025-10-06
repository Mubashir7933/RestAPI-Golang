package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mubashir7933/RestAPI-Golang/internal/config"
	"github.com/Mubashir7933/RestAPI-Golang/internal/http/handlers/student"
	"github.com/Mubashir7933/RestAPI-Golang/internal/storage/sqlite"
)

func main() {
	//load config file
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("Connected to database", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	// setup router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
	// setup Server

	server := http.Server{
		Addr:    cfg.Adress,
		Handler: router,
	}
	// fmt.Printf("Server started on port at port %s", cfg.HTTPServer.Adress)

	slog.Info("Server started on port", slog.String("port", cfg.Adress))

	//Setting up gracefully shutdown by using goroutines and channels
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server: ", err.Error())
		}
	}()

	<-done
	slog.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server Shutdown Failed", slog.String("error", err.Error()))
	}
	slog.Info("Server Exited Properly")
}
