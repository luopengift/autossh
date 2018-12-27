package endpoint

import (
	"fmt"
	"strconv"

	"github.com/luopengift/log"
)

var groupFormat = "%-4v\t%-64s\t%-40v"

// Group group
type Group struct {
	Name      string
	Endpoints Endpoints
}

// Groups groups
type Groups []*Group

// Print PrintGroups
func (grps Groups) Print() {
	log.ConsoleWithGreen(groupFormat, "[ID]", "组名称", "主机数量")
	for idx, grp := range grps {
		log.ConsoleWithGreen(groupFormat, fmt.Sprintf("[%v]", idx), grp.Name, len(grp.Endpoints))
	}
}

// Search search
func (grps Groups) Search(querys ...string) Groups {
	if len(querys) == 1 {
		if id, err := strconv.Atoi(querys[0]); err == nil && id <= grps.Len() {
			return Groups{grps[id]}
		}
	}
	var result Groups
	for _, group := range grps {
		if Find(group.Name, querys...) {
			result = append(result, group)
		}
	}
	return result
}

// Match match
func (grps Groups) Match(match string) Groups {
	var result Groups
	for index, group := range grps {
		if match == strconv.Itoa(index) || match == group.Name {
			result = append(result, group)
		}
	}
	return result
}

// Find checkout name is in Groups.
func (grps Groups) Find(name string) *Group {
	for _, group := range grps {
		if group.Name == name {
			return group
		}
	}
	return nil
}

// Len implements sort.Interface
func (grps Groups) Len() int {
	return len(grps)
}

// Less implements sort.Interface
func (grps Groups) Less(i, j int) bool {
	return grps[i].Endpoints.Len() < grps[j].Endpoints.Len()
}

// Swap implements sort.Interface
func (grps Groups) Swap(i, j int) {
	grps[i], grps[j] = grps[j], grps[i]
}
