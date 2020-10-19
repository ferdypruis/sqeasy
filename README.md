# sqeasy
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ferdypruis/sqeasy)](https://pkg.go.dev/github.com/ferdypruis/sqeasy)

Attempting to make usage of Go's `sql`-package easier to maintain, by preventing you to manually have to 
count and match columns and variables.

## Features
* Map columns to destination values
* Use named parameters

## Examples
### Map columns to destination values
```go
var db *sql.DB


var (
    colA      string
    timestamp time.Time
    count     int
)

columns := sqeasy.SelectColumns{
    `a_column`:    &colA,
    `"timestamp"`: &timestamp,
    `COUNT(*)`:    &count,
}

row := db.QueryRow("SELECT " + columns.ExprList() + " FROM table")
err := columns.Scan(row)
```

### Use named parameters
```go
var db *sql.DB


params := sqeasy.NamedParams{
    "colA":      "notthis",
    "timestamp": time.Now(),
}

query := "SELECT * FROM table WHERE a_column != :colA AND timestamp < :timestamp"
row := sqeasy.QueryRow(db, query, params)
```
