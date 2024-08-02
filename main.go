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

	// Serve static files
	router.PathPrefix("/public/").Handler(staticFileHandler())

	// create view handler
	vh := handlers.NewViewHandler(logger)
	// view router
	vr := router.NewRoute().Subrouter()

	// register view handlers to router
	getv := vr.Methods(http.MethodGet).Subrouter()
	getv.HandleFunc("/", vh.HandleHome)

	// create api handler
	h := handlers.NewHandler(logger, s)
	// api router
	ar := router.NewRoute().Subrouter()

	// regiter book handler to router
	getb := ar.Methods(http.MethodGet).Subrouter()
	getb.HandleFunc("/api/books", h.GetAllBooks)
	getb.HandleFunc("/api/books/{id}", h.GetSingleBook)

	postb := ar.Methods(http.MethodPost).Subrouter()
	postb.HandleFunc("/api/books", h.PostBook)
	postb.Use(m.ValidateBook)

	putb := ar.Methods(http.MethodPut).Subrouter()
	putb.HandleFunc("/api/books/{id}", h.PutBook)
	putb.Use(m.ValidateBook)

	delb := ar.Methods(http.MethodDelete).Subrouter()
	delb.HandleFunc("/api/books/{id}", h.DeleteBook)

	// register author handler to router
	geta := ar.Methods(http.MethodGet).Subrouter()
	geta.HandleFunc("/api/authors", h.GetAllAuthors)
	geta.HandleFunc("/api/authors/{id}", h.GetSingleAuthor)

	posta := ar.Methods(http.MethodPost).Subrouter()
	posta.HandleFunc("/api/authors", h.PostAuthor)
	posta.Use(m.ValidateAuthor)

	puta := ar.Methods(http.MethodPut).Subrouter()
	puta.HandleFunc("/api/authors/{id}", h.PutAuthor)
	puta.Use(m.ValidateAuthor)

	dela := ar.Methods(http.MethodDelete).Subrouter()
	dela.HandleFunc("/api/authors/{id}", h.DeleteAuthor)

	// register genre handler to router
	getg := ar.Methods(http.MethodGet).Subrouter()
	getg.HandleFunc("/api/genres", h.GetAllGenres)
	getg.HandleFunc("/api/genres/{id}", h.GetSingleGenre)

	postg := ar.Methods(http.MethodPost).Subrouter()
	postg.HandleFunc("/api/genres", h.PostGenre)
	postg.Use(m.ValidateGenre)

	putg := ar.Methods(http.MethodPut).Subrouter()
	putg.HandleFunc("/api/genres/{id}", h.PutGenre)
	putg.Use(m.ValidateGenre)

	delg := ar.Methods(http.MethodDelete).Subrouter()
	delg.HandleFunc("/api/genres/{id}", h.DeleteGenre)

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
