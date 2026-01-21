package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(migrationsPath, dsn string) {
	m, err := migrate.New(
		"file://"+migrationsPath,
		dsn,
	)

	if err != nil {
		log.Fatal("migrate.New failed:", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("no new migrations to apply")
		} else {
			log.Fatal("migration failed:", err)
		}
	} else {
		log.Println("migrations applied")
	}

	log.Println("migrations applied")
}
