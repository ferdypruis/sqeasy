package sqeasy_test

import (
	"database/sql"
	"github.com/ferdypruis/sqeasy"
)

func ExampleSelectColumns() {
	var db *sql.DB

	var (
		colA      string
		timestamp sql.NullString
		count     int
	)

	columns := sqeasy.SelectColumns{
		`a_column`:    &colA,
		`"timestamp"`: &timestamp,
		`COUNT(*)`:    &count,
	}

	row := db.QueryRow("SELECT " + columns.ExprList() + " FROM table")
	err := columns.Scan(row)
}
