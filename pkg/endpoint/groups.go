package endpoint

import (
	"fmt"

	"github.com/luopengift/log"
)

var groupFormat = "%-4v\t%-20s\t%-40v"

// Groups groups
type Groups struct {
	Kind   string
	List   []string
	Groups map[string]Endpoints
}

// PrintGroups PrintGroups
func (grps *Groups) PrintGroups() {
	log.ConsoleWithGreen(fmt.Sprintf(groupFormat, "序号", "组名称", "主机数量"))
	for idx, kind := range grps.List {
		log.ConsoleWithGreen(groupFormat, idx, kind, len(grps.Groups[kind]))
	}

}
