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
	"github.com/shahinrahimi/bookbrowse/pkg/author"
	"github.com/shahinrahimi/bookbrowse/pkg/book"
	"github.com/shahinrahimi/bookbrowse/pkg/genre"
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
	if err := s.Init(); err != nil {
		logger.Fatalf("error initilizing DB: %v", err)
	}

	s.MainSeed()

	// create mux serve
	router := mux.NewRouter()

	// create book handler
	bh := book.NewHandler(logger, s)
	// create author handler
	ah := author.NewHandler(logger, s)
	// create genre handler
	gh := genre.NewHandler(logger, s)

	// regiter book handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/books", bh.GetAll)
	router.HandleFunc("/books/{id}", bh.GetSingle)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/books", bh.Create)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/books/{id}", bh.Update)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/books/{id}", bh.Delete)

	// register author handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/authors", ah.GetAll)
	router.HandleFunc("/authors/{id}", ah.GetSingle)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/authors", ah.Create)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/authors/{id}", ah.Update)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/authors/{id}", ah.Delete)

	// register genre handler to router
	router.Methods(http.MethodGet).Subrouter()
	router.HandleFunc("/genres", gh.GetAll)
	router.HandleFunc("/genres/{id}", gh.GetSingle)

	router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/genres", gh.Create)

	router.Methods(http.MethodPut).Subrouter()
	router.HandleFunc("/genres/{id}", gh.Update)

	router.Methods(http.MethodDelete).Subrouter()
	router.HandleFunc("/genres/{id}", gh.Delete)

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
