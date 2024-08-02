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
	"github.com/shahinrahimi/bookbrowse/middlewares"
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

	// s.Seed()

	// create mux serve
	router := mux.NewRouter()
	// create middlware
	m := middlewares.NewMiddlware(logger)
	// add logger middleware
	router.Use(m.Logger)
	// create limitrate middlware
	// rl := middlewares.NewRateLimiter(100, time.Minute*60)
	// router.Use(rl.Limit)

	// create handler
	h := handlers.NewHandler(logger, s)

	// regiter book handler to router
	getb := router.Methods(http.MethodGet).Subrouter()
	getb.HandleFunc("/books", h.GetAllBooks)
	getb.HandleFunc("/books/{id}", h.GetSingleBook)

	postb := router.Methods(http.MethodPost).Subrouter()
	postb.HandleFunc("/books", h.PostBook)
	postb.Use(m.ValidateBook)

	putb := router.Methods(http.MethodPut).Subrouter()
	putb.HandleFunc("/books/{id}", h.PutBook)
	putb.Use(m.ValidateBook)

	delb := router.Methods(http.MethodDelete).Subrouter()
	delb.HandleFunc("/books/{id}", h.DeleteBook)

	// register author handler to router
	geta := router.Methods(http.MethodGet).Subrouter()
	geta.HandleFunc("/authors", h.GetAllAuthors)
	geta.HandleFunc("/authors/{id}", h.GetSingleAuthor)

	posta := router.Methods(http.MethodPost).Subrouter()
	posta.HandleFunc("/authors", h.PostAuthor)
	posta.Use(m.ValidateAuthor)

	puta := router.Methods(http.MethodPut).Subrouter()
	puta.HandleFunc("/authors/{id}", h.PutAuthor)
	puta.Use(m.ValidateAuthor)

	dela := router.Methods(http.MethodDelete).Subrouter()
	dela.HandleFunc("/authors/{id}", h.DeleteAuthor)

	// register genre handler to router
	getg := router.Methods(http.MethodGet).Subrouter()
	getg.HandleFunc("/genres", h.GetAllGenres)
	getg.HandleFunc("/genres/{id}", h.GetSingleGenre)

	postg := router.Methods(http.MethodPost).Subrouter()
	postg.HandleFunc("/genres", h.PostGenre)
	postg.Use(m.ValidateGenre)

	putg := router.Methods(http.MethodPut).Subrouter()
	putg.HandleFunc("/genres/{id}", h.PutGenre)
	putg.Use(m.ValidateGenre)

	delg := router.Methods(http.MethodDelete).Subrouter()
	delg.HandleFunc("/genres/{id}", h.DeleteGenre)

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
