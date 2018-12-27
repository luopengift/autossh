package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/luopengift/autossh/command"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/autossh/console"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

// Run run
func Run(ctx context.Context, conf *config.Config) error {
	var err error
	params := NewParams()

	switch {
	case params.Version:
		fmt.Println(version.VERSION)
		return nil
	case params.Debug:
		log.Warn("params: %v", string(log.Dump(params)))
		log.Warn("config: %v", string(log.Dump(conf)))
		return nil
	case len(os.Args) < 3: //登录交互模式
		if conf.Remote { // 远程获取模式

		} else { // 本地配置模式
			if err = conf.LoadUserConfig(); err != nil {
				return err
			}
			if err = conf.LoadEndpointsConfig(); err != nil {
				return err
			}
			if conf.Script != "" {
				script := strings.Replace(conf.Script, "~", os.Getenv("HOME"), -1)
				b, err := exec.Command(script).CombinedOutput()
				if err != nil {
					return err
				}
				c := &config.Config{} //为了保证user config.endpoints 不被覆盖
				if err = json.Unmarshal(b, c); err != nil {
					return err
				}
				conf.Endpoints = append(conf.Endpoints, c.Endpoints...)
				conf.UseGlobalValues()
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
			endpoint.SetPseudo(params.Pseudo)
			conf.Endpoints = append(conf.Endpoints, endpoint)
		}
		batch := command.NewBatch(params.Fork, params.Timeout)
		return batch.Execute(conf.Endpoints, params.Module, params.Args)
	}
}
