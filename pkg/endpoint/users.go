package endpoint

import (
	"fmt"
	"strconv"

	"github.com/luopengift/log"
)

var userFormat = "%-4v\t%-20s"

// Users users
type Users []string

// Print print users
func (u Users) Print() {
	log.ConsoleWithGreen(fmt.Sprintf(userFormat, "[ID]", "用户名"))
	for idx, user := range u {
		log.ConsoleWithGreen(userFormat, fmt.Sprintf("[%v]", idx), user)
	}
}

// Search search
func (u Users) Search(querys ...string) Users {
	if len(querys) == 1 {
		if id, err := strconv.Atoi(querys[0]); err == nil && id <= u.Len() {
			return Users{u[id]}
		}
	}
	var result Users
	for _, user := range u {
		if FindOr(user, querys...) {
			result = append(result, user)
		}
	}
	return result
}

// Match Match
func (u Users) Match(input string) Users {
	var result Users
	for index, user := range u {
		if input == strconv.Itoa(index) || user == input {
			result = append(result, user)
		}
	}
	return result
}

// Len implements sort.Interface
func (u Users) Len() int {
	return len(u)
}

// Less implements sort.Interface
func (u Users) Less(i, j int) bool {
	return u[i] < u[j]
}

// Swap implements sort.Interface
func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
