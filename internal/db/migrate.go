package db

import (
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(migrationsDir, dsn string) {
	absPath, err := filepath.Abs(migrationsDir)
	if err != nil {
		log.Fatal("failed to get absolute path:", err)
	}

	m, err := migrate.New(
		"file://"+absPath,
		dsn,
	)
	if err != nil {
		log.Fatal("failed to create migrate instance:", err)
	}

	// Up() применяет все миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration failed:", err)
	} else if err == migrate.ErrNoChange {
		log.Println("no new migrations to apply")
	} else {
		log.Println("migrations applied successfully")
	}
}
