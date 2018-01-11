package gorputil

import "strings"

func isSelectQuery(query string) bool {
	if len(query) < len("select") {
		return false
	}
	return strings.ToLower(query[:len("select")]) == "select"
}