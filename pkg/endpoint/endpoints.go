package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
)

// Endpoints ssh.Endpoint slice
type Endpoints []*ssh.Endpoint

// Println println
func (eps Endpoints) Println() {
	format := "%-4v\t%-20s\t%-40s\t%-5s"
	log.ConsoleWithGreen(fmt.Sprintf(format, "序号", "名称", "地址", "用户名"))
	for index, endpoint := range eps {
		users, _ := endpoint.GetUsers()
		log.ConsoleWithGreen(
			fmt.Sprintf(format, index, endpoint.Name, endpoint.Address(), "[ "+strings.Join(users, ", ")+" ]"),
		)
	}
}

// Search search
func (eps Endpoints) Search(search string) Endpoints {
	result := Endpoints{}
	for index, endpoint := range eps {
		if search == strconv.Itoa(index) || strings.Contains(endpoint.Name, search) || strings.Contains(endpoint.Host, search) || strings.Contains(endpoint.IP, search) {
			result = append(result, endpoint)
		}
	}
	eps = result
	return eps
}

// Match match
func (eps Endpoints) Match(match string) Endpoints {
	result := Endpoints{}
	for index, endpoint := range eps {
		if match == strconv.Itoa(index) || match == endpoint.Name || match == endpoint.Host || match == endpoint.IP {
			result = append(result, endpoint)
		}
	}
	eps = result
	return eps
}

// Len implements sort.Interface
func (eps Endpoints) Len() int {
	return len(eps)
}

// Less implements sort.Interface
func (eps Endpoints) Less(i, j int) bool {
	return eps[i].IP < eps[j].IP
}

// Swap implements sort.Interface
func (eps Endpoints) Swap(i, j int) {
	eps[i], eps[j] = eps[j], eps[i]
}
