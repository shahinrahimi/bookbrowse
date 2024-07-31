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
	"github.com/shahinrahimi/bookbrowse/handlers"
	"github.com/shahinrahimi/bookbrowse/stores"
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
	// s := store.NewSqliteStore(logger)
	s := stores.NewSqliteStore(logger)
	defer s.CloseDB()

	// inint store
	if err := s.Init(); err != nil {
		logger.Fatalf("error initilizing DB: %v", err)
	}

	// s.MainSeed()

	// create mux serve
	router := mux.NewRouter()

	// create handler
	h := handlers.NewHandler(logger, s)

	// regiter book handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/books", h.GetAllBooks)
	router.HandleFunc("/books/{id}", h.GetSingleBook)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/books", h.PostBook)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/books/{id}", h.PutBook)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/books/{id}", h.DeleteBook)

	// register author handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/authors", h.GetAllAuthors)
	router.HandleFunc("/authors/{id}", h.GetSingleAuthor)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/authors", h.PostAuthor)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/authors/{id}", h.PutAuthor)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/authors/{id}", h.DeleteAuthor)

	// register genre handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/genres", h.GetAllGenres)
	router.HandleFunc("/genres/{id}", h.GetSingleGenre)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/genres", h.PostGenre)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/genres/{id}", h.PutGenre)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/genres/{id}", h.DeleteGenre)

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
