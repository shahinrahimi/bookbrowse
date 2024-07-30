package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/shahinrahimi/bookbrowse/pkg/book"
	"github.com/shahinrahimi/bookbrowse/store"
)

func main() {
	// create a new custom logger
	logger := log.New(os.Stdout, "[BOOKBROWSE] ", log.LstdFlags)

	// check if .env file
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading .env file: %v", err)
	}
	// load enviromental variable
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		logger.Fatal("error loading enviromental variable")
	}

	// create store
	s := store.NewSqliteStore(logger)
	defer s.CloseDB()

	// inint store
	s.Init()

	// create mux serve
	router := mux.NewRouter()

	// create book handler
	bh := book.NewHandler(logger, s)

	// regiter handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/books", bh.GetAll)
	router.HandleFunc("/books/{id}", bh.GetSingle)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/books", bh.Create)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/books/{id}", bh.Update)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/books/{id}", bh.Delete)

	// craete server
	server := http.Server{
		Addr:     listenAddr,
		Handler:  router,
		ErrorLog: logger,
	}

	go func() {
		logger.Println("Starting server on port", listenAddr)
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalf("error starting server: %s\n", err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	logger.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
	defer cancel()

	fmt.Println("Welcome to BookBrowse API")
}
