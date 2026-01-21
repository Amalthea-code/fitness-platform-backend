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

	r.Post("/auth/register", authHandler.Register)
	r.Post("/auth/login", authHandler.Login)
	r.Post("/auth/logout", authHandler.Logout)
	r.Get("/auth/me", authHandler.Me)

	log.Printf("server started on :%s", cfg.HTTPPort)
	http.ListenAndServe(":"+cfg.HTTPPort, r)
}
