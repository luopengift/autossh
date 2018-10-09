package core

import (
	"os/exec"
	"strings"
)

// Runtime runtime
type Runtime struct {
	User string
}

func getTty() string {
	out, err := exec.Command("/bin/tty").CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func getLoginIP() string {
	out, err := exec.Command("who", "-m").CombinedOutput()
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.Trim(string(out), ")"), "(")[1]
}
