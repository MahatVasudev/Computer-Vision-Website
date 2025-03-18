package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/MahatVasudev/Computer-Vision-Website/server/service/comments"
	"github.com/MahatVasudev/Computer-Vision-Website/server/service/follows"
	"github.com/MahatVasudev/Computer-Vision-Website/server/service/posts"
	"github.com/MahatVasudev/Computer-Vision-Website/server/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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

	c := cors.Handler(
		cors.Options{
			AllowedOrigins: []string{
				"http://localhost:3000",
			}, // Change this to your frontend domain for security
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
			MaxAge:           300,
		},
	)

	router.Use(c)
	fs := http.StripPrefix("/public/", http.FileServer(http.Dir("../public")))
	router.Handle("/public/*", fs)

	// Stores
	userStore := user.NewStore(newapiserver.sqlDB, newapiserver.redDB)
	postStore := posts.NewStore(newapiserver.sqlDB, newapiserver.redDB, "../public/posts/")
	followStore := follows.NewStore(newapiserver.sqlDB, newapiserver.redDB)
	commentStore := comments.NewStore(newapiserver.sqlDB, newapiserver.redDB)

	// Handlers
	userHandler := user.NewHandler(userStore, followStore, postStore, commentStore)
	postHandler := posts.NewHandler(postStore, userStore, commentStore)
	followHandler := follows.NewHandler(followStore, userStore)
	commentHandler := comments.NewHandler(commentStore, userStore, postStore)
	// Routes
	userHandler.RegisterRoutes(`/user/`, router)
	postHandler.RegisterRoutes(`/posts/`, router)
	followHandler.RegisterRoutes("/follow/", router)
	commentHandler.RegisteredRoutes("/comments/", router)

	server := &http.Server{
		Addr:    newapiserver.Addr,
		Handler: router,
	}

	log.Printf("Connecting to Port %s....", newapiserver.Addr)
	return server.ListenAndServe()
}
