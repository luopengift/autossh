package core

import (
	"fmt"
	"os/exec"
	"strings"
)

// Runtime runtime, parse shell cmd `who` or `who -m`
type Runtime struct {
	User string
	Tty  string
	Date string
	IP   string
}

func (r *Runtime) parse(out string) {
	fmt.Println(out)
	out = strings.TrimSpace(out)
	if strings.ContainsAny(out, "()") {
		sList := strings.Split(strings.TrimSuffix(out, ")"), "(")
		out = sList[0]
		r.IP = sList[1]
	} else {
		r.IP = "Localhost"
	}
	sList := strings.Split(out, " ")
	r.User = sList[0]
	r.Tty = sList[1]
	r.Date = strings.Join(sList[2:], " ")
}

func getTty() string {
	out, err := exec.Command("/bin/tty").CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(out)
}

// NewRuntime new runtime
func NewRuntime() (*Runtime, error) {
	out, err := exec.Command("who", "am", "i").CombinedOutput()
	if err != nil {
		return nil, err
	}
	r := &Runtime{}
	r.parse(string(out))
	return r, nil
}
