package mongo

import "strings"

func isNotFound(err error) bool {
	return strings.TrimSpace(err.Error()) == "not found"
}
