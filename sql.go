package namedsql

import (
	"context"
	"database/sql"
)

func QueryContext(ctx context.Context, db *sql.DB, query string, params map[string]interface{}) (*sql.Rows, error) {
	var mapper mapper
	query, mapper = named(query)
	args, err := mapper(params)
	if err != nil {
		return nil, err
	}
	return db.QueryContext(ctx, query, args...)
}

func Query(db *sql.DB, query string, params map[string]interface{}) (*sql.Rows, error) {
	return QueryContext(context.Background(), db, query, params)
}

func QueryRowContext(ctx context.Context, db *sql.DB, query string, params map[string]interface{}) *Row {
	var mapper mapper
	query, mapper = named(query)
	args, err := mapper(params)
	if err != nil {
		return &Row{
			err: err,
		}
	}
	return &Row{
		Row: db.QueryRowContext(ctx, query, args...),
	}
}

func QueryRow(ctx context.Context, db *sql.DB, query string, params map[string]interface{}) *Row {
	return QueryRowContext(context.Background(), db, query, params)
}

func ExecContext(ctx context.Context, db *sql.DB, query string, params map[string]interface{}) (sql.Result, error) {
	var mapper mapper
	query, mapper = named(query)
	args, err := mapper(params)
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}

func Exec(db *sql.DB, query string, params map[string]interface{}) (sql.Result, error) {
	return ExecContext(context.Background(), db, query, params)
}

func PrepareContext(ctx context.Context, db *sql.DB, query string) (*Stmt, error) {
	var mapper mapper
	query, mapper = named(query)

	stmt, err := db.PrepareContext(ctx, query)
	return &Stmt{
		stmt:   stmt,
		mapper: mapper,
	}, err
}

func Prepare(db *sql.DB, query string) (*Stmt, error) {
	return PrepareContext(context.Background(), db, query)
}

type Stmt struct {
	stmt   *sql.Stmt
	mapper mapper
}

func (s *Stmt) Close() error {
	return s.stmt.Close()
}

func (s *Stmt) Exec(args map[string]interface{}) (sql.Result, error) {
	return s.ExecContext(context.Background(), args)
}

func (s *Stmt) ExecContext(ctx context.Context, params map[string]interface{}) (sql.Result, error) {
	args, err := s.mapper(params)
	if err != nil {
		return nil, err
	}
	return s.stmt.ExecContext(ctx, args...)
}

func (s *Stmt) Query(args map[string]interface{}) (*sql.Rows, error) {
	return s.QueryContext(context.Background(), args)
}

func (s *Stmt) QueryContext(ctx context.Context, params map[string]interface{}) (*sql.Rows, error) {
	args, err := s.mapper(params)
	if err != nil {
		return nil, err
	}
	return s.stmt.QueryContext(ctx, args...)
}

func (s *Stmt) QueryRow(args map[string]interface{}) *Row {
	return s.QueryRowContext(context.Background(), args)
}

func (s *Stmt) QueryRowContext(ctx context.Context, params map[string]interface{}) *Row {
	args, err := s.mapper(params)
	if err != nil {
		return &Row{
			err: err,
		}
	}
	row := s.stmt.QueryRowContext(ctx, args...)
	return &Row{
		Row: row,
	}
}

type Row struct {
	*sql.Row
	err error
}

func (r *Row) Scan(dest ...interface{}) error {
	if r.err != nil {
		return r.err
	}
	return r.Row.Scan(dest...)
}
func (r *Row) Err() error {
	if r.err != nil {
		return r.err
	}
	return r.Row.Err()
}
