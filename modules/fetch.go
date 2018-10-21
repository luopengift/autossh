package modules

import (
	"context"

	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
)

// Fetch fetch - Fetches a file from remote nodes
type Fetch struct {
	Src  string `json:"src" yaml:"src"`
	Dest string `json:"dest" yaml:"dest"`
}

// Parse parse
func (mod *Fetch) Parse(cmd string) error {
	args, err := parseArgs(cmd)
	if err != nil {
		return err
	}
	return types.Format(args, mod)
}

// Name name
func (mod *Fetch) Name() string {
	return "fetch"
}

// Run run
func (mod *Fetch) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return nil, endpoint.Download(mod.Src, mod.Dest)
}

func init() {
	ModuleRegister("fetch", &Fetch{})
}
