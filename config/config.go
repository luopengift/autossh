package config

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/types"
)

// Config config
type Config struct {
	Remote  bool            `json:"remote" yaml:"remote"` // 是否开启远程接口获取IP信息
	Shell   bool            `json:"shell" yaml:"shell"`   // 是否支持本地SHELL模式
	Backup  string          `json:"backup" yaml:"backup"` // 是否支持审计模式
	Global  *ssh.Endpoint   `json:"global" yaml:"global"` // global config
	Servers []*ssh.Endpoint `json:"servers" yaml:"servers"`
	result  []*ssh.Endpoint
}

// Init config init
func Init() *Config {
	return &Config{}
}
func (c *Config) LoadConfig(f string) error {
	if err := types.ParseConfigFile(f, c); err != nil {
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
	c.Reset()
	return nil
}

// UseGlobalValues UseGlobalValues
func (c *Config) UseGlobalValues() {
	for _, endpoint := range c.Servers {
		endpoint.Mask(c.Global)
	}
}

// Println println
func (c *Config) Println() {
	format := "%-4v\t%-20s\t%-40s\t%-5s"
	log.ConsoleWithGreen(fmt.Sprintf(format, "序号", "名称", "地址", "用户名"))
	for index, endpoint := range c.result {
		log.ConsoleWithGreen(
			fmt.Sprintf(format, index, endpoint.Name, endpoint.Address(), endpoint.User),
		)
	}
}

// Match match
func (c *Config) Match(match string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for index, endpoint := range c.result {
		if match == strconv.Itoa(index) || match == endpoint.Name || match == endpoint.Host || match == endpoint.IP {
			result = append(result, endpoint)
		}
	}
	c.result = result
	return c.result
}

// Search search
func (c *Config) Search(search string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for index, endpoint := range c.result {
		if search == strconv.Itoa(index) || strings.Contains(endpoint.Name, search) || strings.Contains(endpoint.Host, search) || strings.Contains(endpoint.IP, search) {
			result = append(result, endpoint)
		}
	}
	c.result = result
	return c.result
}

// Reset Reset
func (c *Config) Reset() []*ssh.Endpoint {
	c.result = c.Servers
	return c.result
}

// Add add
func (c *Config) Add(name, host, ip, port, user, password, key string) error {
	endpoint := ssh.NewEndpointWithValue(name, host, ip, port, user, password, key)
	c.Servers = append(c.Servers, endpoint)
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

	c.Servers = append(c.Servers, endpoint)
	c.result = append(c.result, endpoint)
	return nil
}

func (c *Config) Dump(f string) error {
	b, err := types.ToYAML(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f, b, 0644)
}
