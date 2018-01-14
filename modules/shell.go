package modules

import (
	"context"
	"github.com/luopengift/golibs/ssh"
	//"github.com/luopengift/golibs/log"
)

// Execute commands in nodes.
type Shell struct {
	Command string
}

func (mod *Shell) Init(cmd string) error {
	mod.Command = cmd
	return nil
}

func (mod *Shell) Name() string {
	return "shell"
}

func (mod *Shell) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return endpoint.CmdOutBytes(mod.Command)
}

func init() {
	ModuleRegister("shell", &Shell{})
}
