package config

import (
	"io/ioutil"

	"github.com/chzyer/readline"
	"github.com/luopengift/autossh/pkg/endpoint"
	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
)

// Config config
type Config struct {
	Debug              bool          `json:"debug" yaml:"debug"`   // 是否进入调试模式
	Remote             bool          `json:"remote" yaml:"remote"` // 是否开启远程接口获取IP信息
	Shell              bool          `json:"shell" yaml:"shell"`   // 是否支持本地SHELL模式
	Backup             string        `json:"backup" yaml:"backup"` // 是否支持审计模式
	Script             string        `json:"script" yaml:"script"` // 自动更新主机列表脚本, 只有Remote为false时生效
	Global             *ssh.Endpoint `json:"global" yaml:"global"` // global config
	endpoint.Endpoints `json:"endpoints" yaml:"endpoints"`
}

// Init config init
func Init() *Config {
	return &Config{}
}

// LoadConfig loadconfig form a file
func (c *Config) LoadConfig(f string) error {
	if err := types.ParseConfigFile(c, f); err != nil {
		return err
	}
	return nil
}

// LoadRootConfig load rot config
func (c *Config) LoadRootConfig() error {
	if err := c.LoadConfig("/etc/autossh/autossh.yml"); err != nil {
		return err
	}
	return nil
}

// LoadUserConfig load user config
func (c *Config) LoadUserConfig() error {
	if err := c.LoadConfig("~/.autossh/autossh.yml"); err != nil {
		return err
	}
	c.UseGlobalValues()
	return nil
}

// LoadEndpointsConfig load endpoints config
func (c *Config) LoadEndpointsConfig() error {
	newconfig := &Config{} //为了保证user config.endpoints 不被覆盖
	if err := newconfig.LoadConfig("~/.autossh/endpoints.yml"); err != nil {
		return err
	}
	c.Endpoints = append(c.Endpoints, newconfig.Endpoints...)
	c.UseGlobalValues()
	return nil
}

// UseGlobalValues UseGlobalValues
func (c *Config) UseGlobalValues() {
	for _, endpoint := range c.Endpoints {
		endpoint.Mask(c.Global)
	}
}

// Reset Reset
// func (c *Config) Reset() []*ssh.Endpoint {
// 	c.result = c.Endpoints
// 	return c.result
// }

// Add add
func (c *Config) Add(name, host, ip, port, user, password, key string) error {
	endpoint := ssh.NewEndpointWithValue(name, host, ip, port, user, password, key)
	c.Endpoints = append(c.Endpoints, endpoint)
	return nil
}

// ConsoleAdd ConsoleAdd
func (c *Config) ConsoleAdd() error {
	input := ""
	endpoint := ssh.NewEndpoint()
	rl, err := readline.New(readline.StaticPrompt("输入主机名称: "))
	if err != nil {
		return err
	}
	defer rl.Close()
	input, err = rl.Readline()
	if err != nil {
		return err
	}
	endpoint.Name = input

	rl.SetPrompt(readline.StaticPrompt("输入主机地址: "))
	if input, err = rl.Readline(); err != nil {
		return err
	}
	endpoint.Host = input
	// 默认为空,则使用主机名称
	if input == "" {
		endpoint.Host = endpoint.Name
	}

	rl.SetPrompt(readline.StaticPrompt("输入IP地址: "))
	if input, err = rl.Readline(); err != nil {
		return err
	}
	endpoint.IP = input

	rl.SetPrompt(readline.StaticPrompt("输入端口: "))
	if input, err = rl.Readline(); err != nil {
		return err
	}
	endpoint.Port = input

	rl.SetPrompt(readline.StaticPrompt("输入用户名: "))
	if input, err = rl.Readline(); err != nil {
		return err
	}
	endpoint.User = input

	rl.SetPrompt(readline.StaticPrompt("输入密码: "))
	if input, err = rl.Readline(); err != nil {
		return err
	}
	endpoint.Password = input

	rl.SetPrompt(readline.StaticPrompt("输入证书: "))
	if input, err = rl.Readline(); err != nil {
		return err
	}
	endpoint.Key = input

	endpoint.Mask(c.Global)

	c.Endpoints = append(c.Endpoints, endpoint)
	//c.result = append(c.result, endpoint)
	return nil
}

// Dump config to a file
func (c *Config) Dump(f string) error {
	b, err := types.ToYAML(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f, b, 0644)
}
