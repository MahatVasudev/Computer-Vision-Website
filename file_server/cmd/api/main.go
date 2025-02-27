package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type APISERVER struct {
	Addr string
}

func NewApiServer(addr string) *APISERVER {
	return &APISERVER{
		Addr: addr,
	}
}

func (newapi *APISERVER) Run() error {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:5000", "localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
	})

	router.Use(c.Handler)

	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("../public")))
	router.Handle("/public/*", fs)

	server := &http.Server{
		Addr:    newapi.Addr,
		Handler: router,
	}

	log.Printf("Connecting FileServer at port %s....", newapi.Addr)
	return server.ListenAndServe()
}
