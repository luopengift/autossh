package cmd

import (
	"github.com/luopengift/autossh/core"
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/golibs/ssh"
	"github.com/luopengift/types"
	"os"
	"path"
)

func Exec() error {
	var err error
	serverList := &core.ServerList{}
	switch {
	case len(os.Args) == 1: //登录交互模式
		err = types.ParseConfigFile(path.Join(os.Getenv("HOME"), "/.autossh/autossh.yml"), serverList)
		if err != nil {
			return err
		}
		serverList.UseGlobalValues()
		serverList.Reset()
		core.StartConsole(serverList)
	default: //batach模式
		params := NewParams()
		for _, ip := range params.Hosts {
			endpoint := ssh.NewEndpointWithValue("", "", ip, params.Port, params.User, params.Password, params.Key)
			serverList.Servers = append(serverList.Servers, endpoint)
		}
		batch := core.NewBatch(params.Fork, params.Timeout)
		if batch.Execute(serverList.Servers, params.Module, params.Args); err != nil {
			logger.Error("%v", err)
		}
	}
	return nil
}
