package modules

import (
	"context"
	"github.com/luopengift/golibs/ssh"
)

// fetch - Fetches a file from remote nodes
type Fetch struct {
	Src  string `json:"src",yaml:"src"`
	Dest string `json:"dest",yaml:"dest"`
}

func (mod *Fetch) Init(cmd string) error {
	return nil
}

func (mod *Fetch) Name() string {
	return "copy"
}

func (mod *Fetch) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return endpoint.Download(mod.Src, mod.Dest)
}

func init() {
	ModuleRegister("fetch", &Fetch{})
}
