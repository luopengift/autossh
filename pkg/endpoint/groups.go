package endpoint

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luopengift/log"
)

var groupFormat = "%-4v\t%-64s\t%-40v"

// Groups groups
type Groups struct {
	Kind   string
	List   []string
	Groups map[string]Endpoints
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
func (grps *Groups) Search(search ...string) *Groups {
	result := &Groups{
		Kind:   grps.Kind,
		List:   []string{},
		Groups: map[string]Endpoints{},
	}

	for index, group := range grps.List {
		if len(search) == 1 {
			if search[0] == strconv.Itoa(index) {
				result.List = append(result.List, group)
				result.Groups[group] = grps.Groups[group]
				continue
			}
		}
		for _, query := range search {
			if strings.Contains(group, query) {
				result.List = append(result.List, group)
				result.Groups[group] = grps.Groups[group]
				continue
			}
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
