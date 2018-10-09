package modules

import (
	"fmt"
	"strings"
)

func parseArgs(str string) (map[string]string, error) {
	res := make(map[string]string)

	str = strings.Replace(str, " =", "=", -1)
	str = strings.Replace(str, "= ", "=", -1)
	for _, block := range strings.Split(str, " ") {
		values := strings.Split(block, "=")
		if len(values) != 2 {
			return nil, fmt.Errorf("args parse error: %v", values)
		}
		res[strings.TrimSpace(values[0])] = strings.TrimSpace(values[1])
	}
	return res, nil
}
