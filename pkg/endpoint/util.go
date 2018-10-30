package endpoint

import (
	"strconv"
	"strings"

	"github.com/luopengift/ssh"
)

// FindWithIdx find
func FindWithIdx(endpoint *ssh.Endpoint, idx int, querys ...string) bool {
	if len(querys) == 1 {
		if querys[0] == strconv.Itoa(idx) {
			return true
		}
	}
	return endpoint.Find(querys...)
}

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
