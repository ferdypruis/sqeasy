package sqeasy

import (
	"context"
	"database/sql"
)

func QueryContext(ctx context.Context, db *sql.DB, query string, params NamedParams) (*sql.Rows, error) {
	query, args, err := params.Parse(query)
	if err != nil {
		return nil, err
	}
	return db.QueryContext(ctx, query, args...)
}

func Query(db *sql.DB, query string, params NamedParams) (*sql.Rows, error) {
	return QueryContext(context.Background(), db, query, params)
}

func QueryRowContext(ctx context.Context, db *sql.DB, query string, params NamedParams) *Row {
	query, args, err := params.Parse(query)
	if err != nil {
		return &Row{
			err: err,
		}
	}
	return &Row{
		Row: db.QueryRowContext(ctx, query, args...),
	}
}

func QueryRow(db *sql.DB, query string, params NamedParams) *Row {
	return QueryRowContext(context.Background(), db, query, params)
}

func ExecContext(ctx context.Context, db *sql.DB, query string, params NamedParams) (sql.Result, error) {
	query, args, err := params.Parse(query)
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}

func Exec(db *sql.DB, query string, params NamedParams) (sql.Result, error) {
	return ExecContext(context.Background(), db, query, params)
}

func PrepareContext(ctx context.Context, db *sql.DB, query string) (*Stmt, error) {
	query, params := parseNamedQuery(query)
	stmt, err := db.PrepareContext(ctx, query)
	return &Stmt{
		stmt:   stmt,
		params: params,
	}, err
}

func Prepare(db *sql.DB, query string) (*Stmt, error) {
	return PrepareContext(context.Background(), db, query)
}

type Stmt struct {
	stmt   *sql.Stmt
	params []string
}

func (s *Stmt) Close() error {
	return s.stmt.Close()
}

func (s *Stmt) ExecContext(ctx context.Context, named NamedParams) (sql.Result, error) {
	args, err := named.Args(s.params)
	if err != nil {
		return nil, err
	}
	return s.stmt.ExecContext(ctx, args...)
}

func (s *Stmt) Exec(args NamedParams) (sql.Result, error) {
	return s.ExecContext(context.Background(), args)
}

func (s *Stmt) QueryContext(ctx context.Context, named NamedParams) (*sql.Rows, error) {
	args, err := named.Args(s.params)
	if err != nil {
		return nil, err
	}
	return s.stmt.QueryContext(ctx, args...)
}

func (s *Stmt) Query(args NamedParams) (*sql.Rows, error) {
	return s.QueryContext(context.Background(), args)
}

func (s *Stmt) QueryRowContext(ctx context.Context, named NamedParams) *Row {
	args, err := named.Args(s.params)
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

func (s *Stmt) QueryRow(args NamedParams) *Row {
	return s.QueryRowContext(context.Background(), args)
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
