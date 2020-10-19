package sqeasy_test

import (
	"database/sql"
	"github.com/ferdypruis/sqeasy"
	"time"
)

func ExampleNamedParams() {
	var db *sql.DB

	params := sqeasy.NamedParams{
		"colA":      "notthis",
		"timestamp": time.Now(),
	}

	query := "SELECT * FROM table WHERE a_column != :colA AND timestamp < :timestamp"
	row := sqeasy.QueryRow(db, query, params)
}
