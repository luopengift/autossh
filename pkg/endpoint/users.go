package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luopengift/log"
)

var userFormat = "%-4v\t%-20s"

// Users users
type Users []string

// Print print users
func (s Users) Print() {
	log.ConsoleWithGreen(fmt.Sprintf(userFormat, "[ID]", "用户名"))
	for idx, user := range s {
		log.ConsoleWithGreen(userFormat, fmt.Sprintf("[%v]", idx), user)
	}
}

// Search search
func (s Users) Search(querys ...string) Users {
	var result Users
	for index, user := range s {
		if len(querys) == 1 {
			if querys[0] == strconv.Itoa(index) {
				result = append(result, user)
				continue
			}
		}
		for _, query := range querys {
			if strings.Contains(user, query) {
				result = append(result, user)
				continue
			}
		}
	}
	return result
}

// Match Match
func (s Users) Match(input string) Users {
	var result Users
	for index, user := range s {
		if input == strconv.Itoa(index) || user == input {
			result = append(result, user)
		}
	}
	return result
}
