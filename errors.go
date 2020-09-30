package namedsql

import (
	"fmt"
)

type ParameterCountError struct {
	Expected []string
	Got      map[string]interface{}
}

func (e ParameterCountError) Error() string {
	return fmt.Sprintf("namedsql: expected %d parameters, got %d", len(e.Expected), len(e.Got))
}

type MissingParameterError struct {
	Missing string
	Got     map[string]interface{}
}

func (e MissingParameterError) Error() string {
	return fmt.Sprintf("namedsql: expected parameter %q", e.Missing)
}
