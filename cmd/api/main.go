package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"

	"github.com/dmitrijkrasikov/fitness-platform-backend/internal/config"
	"github.com/dmitrijkrasikov/fitness-platform-backend/internal/db"
	"github.com/dmitrijkrasikov/fitness-platform-backend/internal/handlers"
)

func main() {
	cfg := config.Load()

	absPath, err := filepath.Abs("./migrations")
	if err != nil {
		log.Fatal("failed to get absolute path:", err)
	}
	database, err := db.NewPostgres(cfg)
	db.RunMigrations(absPath, cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	authHandler := &handlers.AuthHandler{DB: database}
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register) // POST /auth/register
		r.Post("/login", authHandler.Login)       // POST /auth/login
		r.Post("/logout", authHandler.Logout)     // POST /auth/logout
		r.Get("/me", authHandler.Me)              // POST /auth/me
	})

	usersHandler := &handlers.UsersHandler{DB: database}
	r.Route("/users", func(r chi.Router) {
		r.Get("/", usersHandler.ListUsers)   // GET /users
		r.Post("/", usersHandler.CreateUser) // POST /users
		r.Get("/{id}", usersHandler.GetUser) // GET /users/:id
	})

	log.Printf("server started on :%s", cfg.HTTPPort)
	http.ListenAndServe(":"+cfg.HTTPPort, r)
}
