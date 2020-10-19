package sqeasy

import "strings"

// SelectColumns is a mapping of expression, like a column, to a destination variable
type SelectColumns map[string]interface{}

// ExprList returns all the expressions separated by comma
// Use this in your SELECT statement
func (sc SelectColumns) ExprList() string {
	exprlist := make([]string, len(sc))
	var i int
	for expr := range sc {
		exprlist[i] = expr
		i++
	}
	return strings.Join(exprlist, `,`)
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
