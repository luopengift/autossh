package runtime

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/luopengift/autossh/pkg/endpoint"
)

// Runtime runtime, parse shell cmd `who` or `who -m`
type Runtime struct {
	Super     bool //超级模式
	User      string
	Tty       string
	Date      string
	IP        string
	Endpoints endpoint.Endpoints
	Groups    *endpoint.Groups
}

func (r *Runtime) String() string {
	var ps string
	if r.Super {
		ps = "+"
	} else {
		ps = ">"
	}
	return fmt.Sprintf("[%s]%s ", r.User, ps)
}

// SetEndpoints set endpoints
func (r *Runtime) SetEndpoints(eps endpoint.Endpoints) {
	r.Endpoints = eps
	r.SyncGroups("Group")
}

// SyncGroups sync groups
func (r *Runtime) SyncGroups(kind string) {
	r.Groups = r.Endpoints.Groups(kind)
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
