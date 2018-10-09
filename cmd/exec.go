package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/luopengift/autossh/config"
	"github.com/luopengift/autossh/core"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
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
			err = types.ParseConfigFile("~/.autossh/autossh.yml", conf)
			if err != nil {
				log.Error("%v", err)
				return err
			}
			conf.UseGlobalValues()
			conf.Reset()
		}
		return core.StartConsole(ctx, conf)
	default: //batach模式
		hosts, err := params.Hosts()
		if err != nil {
			return err
		}
		for _, ip := range hosts {
			endpoint := ssh.NewEndpointWithValue("", "", ip, params.Port, params.User, params.Password, params.Key)
			conf.Servers = append(conf.Servers, endpoint)
		}
		batch := core.NewBatch(params.Fork, params.Timeout)
		return batch.Execute(conf.Servers, params.Module, params.Args)
	}
}
