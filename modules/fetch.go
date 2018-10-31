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
func (s *Fetch) Parse(cmd string) error {
	args, err := parseArgs(cmd)
	if err != nil {
		return err
	}
	return types.Format(args, s)
}

// Name name
func (s *Fetch) Name() string {
	return "fetch"
}

// Run run
func (s *Fetch) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return nil, endpoint.Download(s.Src, s.Dest)
}

// Close endpoint
func (s *Fetch) Close(endpoint *ssh.Endpoint) error {
	return endpoint.Close()
}
func init() {
	ModuleRegister("fetch", &Fetch{})
}
