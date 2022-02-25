# sqeasy
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ferdypruis/sqeasy)](https://pkg.go.dev/github.com/ferdypruis/sqeasy)

Small wrappers to make usage of Go's `sql`-package with a CockroachDB database easier to maintain, by for example
mapping variables to positional arguments.

## Examples
### Map columns to destination values
Instead of summing all destination variables in `Scan()`, map the columns onto your variables with `sqeasy.SelectColumns`
and have them scanned into them.
```go
var db *sql.DB

var (
    colA      string
    timestamp time.Time
    count     int
)

columns := sqeasy.SelectColumns{
    {`a_column`,    &colA},
    {`"timestamp"`, &timestamp},
    {`COUNT(*)`,    &count},
}

row := db.QueryRow("SELECT " + columns.ExprList() + " FROM table")
err := columns.Scan(row)
```

### Use named parameters
Instead of the $1, $2, $3 etc positional parameters in your queries, use named ones. Bind values to the names 
using `sqeasy.NamedParams`.
```go
var db *sql.DB

params := sqeasy.NamedParams{
    {"colA",      "notthis"},
    {"timestamp", time.Now()},
}

query := "SELECT * FROM table WHERE a_column != :colA AND timestamp < :timestamp"
row := sqeasy.QueryRow(db, query, params)
```

The same `sqeasy.NamedParams` could be used to for example generate INSERT statements.
```go
var db *sql.DB

params := sqeasy.NamedParams{
    "a_column":  "this",
    "timestamp": time.Now(),
}

query := "INSERT INTO table (" + params.ExprList() + ") VALUES (" + params.Params() + ")"
row := sqeasy.QueryRow(db, query, params)
```
