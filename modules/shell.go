package modules

import (
	"context"

	"github.com/luopengift/ssh"
)

// Shell Execute commands in nodes.
type Shell struct {
	Command string
}

// Init init
func (mod *Shell) Init(cmd string) error {
	mod.Command = cmd
	return nil
}

// Name name
func (mod *Shell) Name() string {
	return "shell"
}

// Run run module
func (mod *Shell) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return endpoint.CmdOutBytes(mod.Command)
}

func init() {
	ModuleRegister("shell", &Shell{})
}
