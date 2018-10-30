package endpoint

import (
	"fmt"
	"strconv"

	"github.com/luopengift/log"
)

var groupFormat = "%-4v\t%-64s\t%-40v"

// Groups groups
type Groups struct {
	Kind   string
	List   []string
	Groups map[string]Endpoints
}

// Endpoints return the list of groups
func (grps *Groups) Endpoints() Endpoints {
	var result Endpoints
	for _, endpoints := range grps.Groups {
		result = append(result, endpoints...)
	}
	return result
}

// Len len
func (grps *Groups) Len() int {
	return len(grps.List)
}

// Print PrintGroups
func (grps *Groups) Print() {
	log.ConsoleWithGreen(groupFormat, "[ID]", "组名称", "主机数量")
	for idx, kind := range grps.List {
		log.ConsoleWithGreen(groupFormat, fmt.Sprintf("[%v]", idx), kind, len(grps.Groups[kind]))
	}
}

// Search search
func (grps *Groups) Search(querys ...string) *Groups {
	result := &Groups{
		Kind:   grps.Kind,
		List:   []string{},
		Groups: map[string]Endpoints{},
	}

	for index, group := range grps.List {
		if len(querys) == 1 {
			if querys[0] == strconv.Itoa(index) {
				result.List = append(result.List, group)
				result.Groups[group] = grps.Groups[group]
				continue
			}
		}

		if Find(group, querys...) {
			result.List = append(result.List, group)
			result.Groups[group] = grps.Groups[group]
		}
	}
	return result
}

// Match match
func (grps *Groups) Match(match string) *Groups {
	result := &Groups{
		Kind:   grps.Kind,
		List:   []string{},
		Groups: map[string]Endpoints{},
	}
	for index, group := range grps.List {
		if match == strconv.Itoa(index) || match == group {
			result.List = append(result.List, group)
			result.Groups[group] = grps.Groups[group]
		}
	}
	return result
}
