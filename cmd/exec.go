package cmd

import (
	"fmt"
	"os"

	"github.com/luopengift/autossh/core"
	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
	"github.com/luopengift/version"
)

// Exec exec
func Exec() error {
	var err error
	serverList := &core.ServerList{}
	params := NewParams()
	if params.Version {
		fmt.Println(version.VERSION)
		return nil
	}
	switch {
	case len(os.Args) < 3: //登录交互模式
		err = types.ParseConfigFile("~/.autossh/autossh.yml", serverList)
		if err != nil {
			return err
		}
		serverList.UseGlobalValues()
		serverList.Reset()
		return core.StartConsole(serverList)
	default: //batach模式
		hosts, err := params.Hosts()
		if err != nil {
			return err
		}
		for _, ip := range hosts {
			endpoint := ssh.NewEndpointWithValue("", "", ip, params.Port, params.User, params.Password, params.Key)
			serverList.Servers = append(serverList.Servers, endpoint)
		}
		batch := core.NewBatch(params.Fork, params.Timeout)
		return batch.Execute(serverList.Servers, params.Module, params.Args)
	}
}
