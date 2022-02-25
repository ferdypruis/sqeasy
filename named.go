package sqeasy

import (
	"fmt"
)

// NamedParams enables the use of named parameters with any database by converting the query to use positional arguments
type NamedParams map[string]interface{}

// Parse replaces the named parameters with positional parameters and returns a slice of corresponding values
func (np NamedParams) Parse(query string) (string, []interface{}, error) {
	query, params := parseNamedQuery(query)
	args, err := np.Args(params)

	return query, args, err
}

// args returns the value for each parameter on the position as indicated in the param slice
func (np NamedParams) Args(params []string) ([]interface{}, error) {
	args := make([]interface{}, len(params))

	// Make sure every argument has been provided
	// TODO This prevents being able to use a param twice
	if len(params) != len(np) {
		return args, ParameterCountError{Expected: params, Got: np}
	}

	var ok bool
	for i, name := range params {
		if args[i], ok = np[name]; !ok {
			return args, MissingParameterError{Missing: name, Got: np}
		}
	}

	return args, nil
}

func parseNamedQuery(q string) (query string, params []string) {
	const tokenStart = ':'

	l := len(q)
	var i, offset int
	for {
		// Scan for tokenStart
		for i < l && q[i] != tokenStart {
			i++
			continue
		}

		// EOL
		if i == l {
			break
		}

		// Add part up to tokenStart to output
		query += q[offset:i]
		i++
		offset = i

		// If this is the last char, or the next is also a tokenStart, take it literal
		if i == l || q[i] == tokenStart {
			query += string(tokenStart)
			i++
			offset = i
			continue
		}

		// Scan for the end of the name
		for i < l {
			if q[i] >= 'a' && q[i] <= 'z' {
				i++
				continue
			}
			if q[i] >= 'A' && q[i] <= 'Z' {
				i++
				continue
			}
			if q[i] >= '0' && q[i] <= '9' {
				i++
				continue
			}
			if q[i] == '_' {
				i++
				continue
			}

			break
		}

		// Register the token
		params = append(params, q[offset:i])
		// Add positional argument of token to query
		query += fmt.Sprintf("$%d", len(params))
		offset = i

		continue
	}

	// Add remainder
	query += q[offset:]

	return
}

type ParameterCountError struct {
	Expected []string
	Got      NamedParams
}

func (e ParameterCountError) Error() string {
	return fmt.Sprintf("sqeasy: expected %d parameters, got %d", len(e.Expected), len(e.Got))
}

type MissingParameterError struct {
	Missing string
	Got     NamedParams
}

func (e MissingParameterError) Error() string {
	return fmt.Sprintf("sqeasy: expected parameter %q", e.Missing)
}
