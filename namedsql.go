package namedsql

import (
	"fmt"
)

type mapper func(params map[string]interface{}) ([]interface{}, error)

// named replaces the named parameters with positional parameters and creates the mapper to map data into them
func named(query string) (string, mapper) {
	query, names := parseQuery(query)
	mapper := func(params map[string]interface{}) ([]interface{}, error) {
		return args(names, params)
	}

	return query, mapper
}

func parseQuery(q string) (query string, names []string) {
	const param = ':'

	l := len(q)
	var i, offset int
	for {
		// Scan for param
		for i < l && q[i] != param {
			i++
			continue
		}

		// EOL
		if i == l {
			break
		}

		// Add part up to param to output
		query += q[offset:i]
		i++
		offset = i

		// If this is the last char, or the next is also a param, take it literal
		if i == l || q[i] == param {
			query += string(param)
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

		names = append(names, q[offset:i])
		query += fmt.Sprintf("$%d", len(names))
		offset = i

		continue
	}

	query += q[offset:]

	return
}

// args maps the parameters to the correct position in the argument list
func args(names []string, params map[string]interface{}) ([]interface{}, error) {
	args := make([]interface{}, len(names))

	// TODO This prevents being able to use a param twice
	if len(params) != len(names) {
		return args, ParameterCountError{Expected: names, Got: params}
	}

	var ok bool
	for i, name := range names {
		if args[i], ok = params[name]; !ok {
			return args, MissingParameterError{Missing: name, Got: params}
		}
	}

	return args, nil
}
