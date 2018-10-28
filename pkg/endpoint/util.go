package endpoint

import (
	"strconv"

	"github.com/luopengift/ssh"
)

func FindWithIdx(endpoint *ssh.Endpoint, idx int, querys ...string) bool {
	if len(querys) == 1 {
		if querys[0] == strconv.Itoa(idx) {
			return true
		}
	}
	return endpoint.Find(querys...)
}
