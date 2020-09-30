# namedsql
Use named parameters with Go's database/sql

## Examples
```go
query := "UPDATE t SET colA = :colA, colB = :col_b WHERE colC = :c"
params := map[string]interface{}{
    "colA":  time.Now(),
    "col_b": "fubar",
    "c":  12,
}
result, _ := namedsql.Exec(db, query, params)
```

```go
stmt, _ := namedsql.Prepare(db, "UPDATE t SET colA = :colA, colB = :col_b WHERE colC = :c")
result, _ := stmt.Exec(map[string]interface{}{
    "colA":  time.Now(),
    "col_b": "fubar",
    "c":  12,
})
```