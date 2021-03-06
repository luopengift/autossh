package modules

import (
	"context"
	"path"

	"github.com/luopengift/ssh"
)

// Script script module
type Script struct {
	Path string `json:"path" yaml:"path"`
}

// Parse parse
func (s *Script) Parse(cmd string) error {
	s.Path = cmd
	return nil
}

// Name name
func (s *Script) Name() string {
	return "srcipt"
}

// Run run module
func (s *Script) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	filepath := path.Join("/tmp/.autossh", path.Base(s.Path))
	if err := endpoint.Upload(s.Path, filepath, 0755); err != nil {
		return nil, err
	}
	return endpoint.CmdOutBytes(filepath)
}

// Close endpoint
func (s *Script) Close(endpoint *ssh.Endpoint) error {
	return endpoint.Close()
}
func init() {
	ModuleRegister("script", &Script{})
}
