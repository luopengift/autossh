package modules

import (
	"context"
	"fmt"

	"github.com/luopengift/ssh"
)

// Moduler
type Moduler interface {
	Name() string
	Init(command string) error
	Run(ctx context.Context, endpoint *ssh.Endpoint) ([]byte, error)
}

var Modules = make(map[string]Moduler)

func ModuleRegister(name string, module Moduler) error {
	if _, ok := Modules[name]; ok {
		return fmt.Errorf("module is exist! %v", name)
	}
	Modules[name] = module
	return nil
}

/*
func Init() {
	Modules = make(map[string]Moduler)
}
*/
