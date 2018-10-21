package modules

import (
	"context"

	"github.com/luopengift/ssh"
)

// Shell Execute commands in nodes.
type Shell struct {
	Command string `json:"command" yaml:"commad"`
}

// Parse parse
func (s *Shell) Parse(cmd string) error {
	s.Command = cmd
	return nil
}

// Name name
func (s *Shell) Name() string {
	return "shell"
}

// Run run module
func (s *Shell) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return endpoint.CmdOutBytes(s.Command)
}

func init() {
	ModuleRegister("shell", &Shell{})
}
