package sqeasy

import (
	"database/sql"
	"time"
)

func ExampleArg() {
	var db *sql.DB

	var args []interface{}
	query := `SELECT * FROM table WHERE a_column = ` + Arg(&args, `fubar`) + ` AND timestamp < ` + Arg(&args, time.Now())

	_ = db.QueryRow(query, args...)
}
