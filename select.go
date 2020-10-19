package sqeasy

import "strings"

// SelectColumns is a mapping of expression, like a column, to a destination variable
type SelectColumns map[string]interface{}

// exprs returns the slice of expressions
func (sc SelectColumns) exprs() []string {
	exprs := make([]string, len(sc))
	var i int
	for expr := range sc {
		exprs[i] = expr
		i++
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
	var i int
	for _, dest := range sc {
		dests[i] = dest
		i++
	}
	return row.Scan(dests...)
}
