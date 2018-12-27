package endpoint

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
)

var endpointFormat = "%-3v\t%-64s\t%-40s\t%-20s\t%-10s"

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

// Groups format Endpoints to Groups
func (eps Endpoints) Groups(kind string) Groups {
	groups := Groups{&Group{
		Name:      "_",
		Endpoints: Endpoints{},
	}}
	for _, endpoint := range eps {
		if value, ok := endpoint.Labels[kind]; ok {
			if group := groups.Find(value); group != nil {
				group.Endpoints = append(group.Endpoints, endpoint)
			} else {
				groups = append(groups, &Group{
					Name:      value,
					Endpoints: Endpoints{endpoint},
				})
			}
		} else {
			groups = append(groups, &Group{
				Name:      "_",
				Endpoints: Endpoints{endpoint},
			})
		}
	}
	sort.Sort(groups)
	return groups
}

// Search search
func (eps Endpoints) Search(querys ...string) Endpoints {
	if len(querys) == 1 {
		if id, err := strconv.Atoi(querys[0]); err == nil && id <= eps.Len() {
			return Endpoints{eps[id]}
		}
	}
	var result Endpoints
	for _, endpoint := range eps {
		if Find(endpoint.Name, querys...) || Find(endpoint.Host, querys...) || Find(endpoint.IP, querys...) {
			result = append(result, endpoint)
		}
	}
	return result
}

// Match match
func (eps Endpoints) Match(match string) Endpoints {
	result := Endpoints{}
	for _, endpoint := range eps {
		if match == endpoint.Name || match == endpoint.Host || match == endpoint.IP {
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

// Close close endpoints
func (eps Endpoints) Close() error {
	for _, ep := range eps {
		ep.Close()
	}
	return nil
}
