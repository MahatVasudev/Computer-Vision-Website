package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MahatVasudev/Computer-Vision-Website/server/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redis/go-redis/v9"
)

type APIServer struct {
	Addr  string
	sqlDB *sql.DB
	redDB *redis.Client
}

func NewAPIServer(addr string, sqlDB *sql.DB, redDB *redis.Client) *APIServer {
	return &APIServer{
		Addr:  addr,
		sqlDB: sqlDB,
		redDB: redDB,
	}
}

func (newapiserver *APIServer) Run() error {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("../public")))
	router.Handle("/public/*", fs)

	userStore := user.NewStore(newapiserver.sqlDB, newapiserver.redDB)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(`/user/`, router)
	server := &http.Server{
		Addr:    newapiserver.Addr,
		Handler: router,
	}

	log.Printf("Connecting to Port %s....", newapiserver.Addr)
	return server.ListenAndServe()
}
