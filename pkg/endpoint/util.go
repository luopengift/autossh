package endpoint

import (
	"strings"
)

// FindOr find key is contains in one of querys.
func FindOr(key string, querys ...string) bool {
	for _, query := range querys {
		if strings.Contains(key, query) {
			return true
		}
	}
	return false
}

// Find key must contains all querys!
func Find(key string, querys ...string) bool {
	for _, query := range querys {
		if !FindOr(key, query) {
			return false
		}
	}
	return true
}
