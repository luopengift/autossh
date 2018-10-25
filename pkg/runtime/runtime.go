package runtime

import (
	"fmt"
	"os/user"

	"github.com/luopengift/autossh/pkg/endpoint"
)

// Runtime runtime, parse shell cmd `who` or `who -m`
type Runtime struct {
	*user.User
	Super     bool //超级模式
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
	return fmt.Sprintf("[%s]%s ", r.User.Username, ps)
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

// NewRuntime new runtime
func NewRuntime() (*Runtime, error) {
	var err error
	r := &Runtime{}
	r.User, err = user.Current()
	if err != nil {
		return nil, err
	}
	return r, nil
}
