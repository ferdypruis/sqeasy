# sqeasy
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ferdypruis/sqeasy)](https://pkg.go.dev/github.com/ferdypruis/sqeasy)

Attempting to make usage of Go's `sql`-package easier to maintain, by preventing you to manually have to 
count and match columns and variables.

## Features
* Map columns to destination values
* Use named parameters

## Example
Instead of;
```go
var (
   colA      string
   timestamp time.Time
   count     int
)

query := `SELECT a_column, "timestamp", COUNT(*) FROM table WHERE a_column != $1 AND timestamp < $2`
err := db.QueryRow(query, "notthis", time.Now()).Scan(
	&colA,
	&timestamp,
	&count,
)
```

You can do it sqeasy with;
```go
var (
   colA      string
   timestamp time.Time
   count     int
)

// Map columns to destination values
columns := sqeasy.SelectColumns{
    {`a_column`,    &colA},
    {`"timestamp"`, &timestamp},
    {`COUNT(*)`,    &count},
}

query := `SELECT ` + columns.ExprList() + ` FROM table WHERE a_column != :colA AND timestamp < :timestamp`

// Use named parameters
params := sqeasy.NamedParams{
    "colA":      "notthis",
    "timestamp": time.Now(),
}
err := columns.Scan(sqeasy.QueryRow(db, query, params))
```
