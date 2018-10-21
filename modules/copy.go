package modules

import (
	"context"
	"os"

	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
)

//Copy copies files to remote locations
type Copy struct {
	Src  string      `json:"src" yaml:"src"`
	Dest string      `json:"dest" yaml:"dest"`
	Mode os.FileMode `json:"mode" yaml:"mode"`
}

func NewCopy() *Copy {
	return &Copy{
		Mode: 0644,
	}
}

// Name name
func (mod *Copy) Name() string {
	return "copy"
}

// Parse parse
func (s *Copy) Parse(cmd string) error {
	args, err := parseArgs(cmd)
	if err != nil {
		return err
	}
	return types.Format(args, s)
}

// Run run
func (s *Copy) Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error) {
	return nil, endpoint.Upload(s.Src, s.Dest, 0644)
}

func init() {
	ModuleRegister("copy", NewCopy())
}
