package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

// RunMigrations applies database migrations.
func RunMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	// Replace "migrations" with the path to your migrations folder
	if err := goose.Up(db, "./migrations"); err != nil {
		return err
	}
	log.Println("Database migrations applied successfully.")
	return nil
}
