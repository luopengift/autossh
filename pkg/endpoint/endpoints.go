package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
)

var endpointFormat = "%-3v\t%-20s\t%-40s\t%-20s\t%-10s"

// Endpoints ssh.Endpoint slice
type Endpoints []*ssh.Endpoint

// Print PrintEndpoints
func (eps Endpoints) Print() {
	log.ConsoleWithGreen(endpointFormat, "[ID]", "名称", "地址", "用户名", "组")
	for idx, endpoint := range eps {
		users, _ := endpoint.GetUsers()
		log.ConsoleWithGreen(endpointFormat, fmt.Sprintf("[%v]", idx), endpoint.Name, endpoint.Address(), fmt.Sprintf("[%v]", strings.Join(users, ", ")), endpoint.Labels["Group"])
	}
}

// Groups groups
func (eps Endpoints) Groups(kind string) *Groups {
	set := make(map[string]struct{})
	g := &Groups{
		Kind:   kind,
		List:   []string{},
		Groups: map[string]Endpoints{},
	}
	g.Kind = kind
	for _, endpoint := range eps {
		if grp, ok := endpoint.Labels[kind]; ok {
			g.Groups[grp] = append(g.Groups[grp], endpoint)
			set[grp] = struct{}{}
		}
	}
	for key := range set {
		g.List = append(g.List, key)
	}
	return g
}

// Search search
func (eps Endpoints) Search(search string) Endpoints {
	result := Endpoints{}
	for index, endpoint := range eps {
		if search == strconv.Itoa(index) || strings.Contains(endpoint.Name, search) || strings.Contains(endpoint.Host, search) || strings.Contains(endpoint.IP, search) {
			result = append(result, endpoint)
		}
	}
	return result
}

// Match match
func (eps Endpoints) Match(match string) Endpoints {
	result := Endpoints{}
	for index, endpoint := range eps {
		if match == strconv.Itoa(index) || match == endpoint.Name || match == endpoint.Host || match == endpoint.IP {
			result = append(result, endpoint)
		}
	}
	return result
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
