package sqeasy

import "strings"

// SelectColumns is a mapping of expression, like a column, to a destination variable
type SelectColumns []struct {
	Expr string
	Dest interface{}
}

// exprs returns the slice of expressions
func (sc SelectColumns) exprs() []string {
	exprs := make([]string, len(sc))
	for i, s := range sc {
		exprs[i] = s.Expr
	}

	return exprs
}

// ExprList returns all the expressions separated by comma
// Use this in your SELECT statement
func (sc SelectColumns) ExprList() string {
	return strings.Join(sc.exprs(), `,`)
}

// Scan copies the columns from the row into the values
func (sc SelectColumns) Scan(row interface {
	Scan(dest ...interface{}) error
}) error {
	dests := make([]interface{}, len(sc))
	for i, s := range sc {
		dests[i] = s.Dest
	}
	return row.Scan(dests...)
}
