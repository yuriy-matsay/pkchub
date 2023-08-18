package sqlite

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

func migrationsUp(db *sql.DB) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://sqlite/schema",
		"sqlite", driver)
	m.Up()
	return err
}
