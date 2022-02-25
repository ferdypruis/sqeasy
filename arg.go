package sqeasy

import (
	"strconv"
)

// Arg adds a value to args, returning it's position
func Arg(args *[]interface{}, arg interface{}) string {
	*args = append(*args, arg)
	return `$` + strconv.Itoa(len(*args))
}
