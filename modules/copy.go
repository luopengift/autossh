package modules

import (
	"context"

	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
)

//Copy copies files to remote locations
type Copy struct {
	Src  string `json:"src" yaml:"src"`
	Dest string `json:"dest" yaml:"dest"`
}

// Name name
func (mod *Copy) Name() string {
	return "copy"
}

// Init init
func (mod *Copy) Init(cmd string) error {
	args, err := parseArgs(cmd)
	if err != nil {
		return err
	}
	return types.Format(args, mod)
}

// Run run
func (mod *Copy) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return nil, endpoint.Upload(mod.Src, mod.Dest)
}

func init() {
	ModuleRegister("copy", new(Copy))
}
