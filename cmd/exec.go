package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/luopengift/autossh/command"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/autossh/console"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

// Exec exec
func Exec(ctx context.Context, conf *config.Config) error {
	var err error
	params := NewParams()
	if params.Version {
		fmt.Println(version.VERSION)
		return nil
	}
	switch {
	case len(os.Args) < 3: //登录交互模式
		if conf.Remote { // 远程获取模式

		} else { // 本地配置模式
			if err = conf.LoadUserConfig(); err != nil {
				return err
			}
		}
		return console.StartConsole(ctx, conf)
	default: //batach模式
		hosts, err := params.Hosts()
		if err != nil {
			return err
		}
		for _, ip := range hosts {
			endpoint := ssh.NewEndpointWithValue("", "", ip, params.Port, params.User, params.Password, params.Key)
			conf.Endpoints = append(conf.Endpoints, endpoint)
		}
		batch := command.NewBatch(params.Fork, params.Timeout)
		return batch.Execute(conf.Endpoints, params.Module, params.Args)
	}
}
