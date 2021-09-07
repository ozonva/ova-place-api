package health

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// NewDatabaseCheck returns Service.
func NewDatabaseCheck(db *sqlx.DB) Service {
	return &database{
		db: *db,
	}
}

// database is a Service implementation.
type database struct {
	db sqlx.DB
}

// Health checks database status and returns an error if the database has some problems.
func (d *database) Health() error {
	_, err := d.db.Exec("SELECT 1")
	if err != nil {
		return fmt.Errorf("cannot Get: %w", err)
	}

	return nil
}
