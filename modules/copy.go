package modules

import (
	"context"
	"github.com/luopengift/golibs/ssh"
	"github.com/luopengift/types"
)

//Copies files to remote locations
type Copy struct {
	Src  string `json:"src",yaml:"src"`
	Dest string `json:"dest",yaml:"dest"`
}

func (mod *Copy) Name() string {
	return "copy"
}

func (mod *Copy) Init(cmd string) error {
	args, err := parseArgs(cmd)
	if err != nil {
		return err
	}
	return types.Format(args, mod)
}

func (mod *Copy) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return endpoint.Upload(mod.Src, mod.Dest)
}

func init() {
	ModuleRegister("copy", new(Copy))
}
