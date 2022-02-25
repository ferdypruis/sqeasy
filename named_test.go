package sqeasy_test

import (
	"database/sql"
	"time"

	"github.com/ferdypruis/sqeasy"
)

func ExampleNamedParams() {
	var db *sql.DB

	params := sqeasy.NamedParams{
		"colA":      "notthis",
		"timestamp": time.Now(),
	}

	query := "SELECT * FROM table WHERE a_column != :colA AND timestamp < :timestamp"
	_ = sqeasy.QueryRow(db, query, params)
}
