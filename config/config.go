package config

import (
	"fmt"
	"strconv"
	"strings"

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
	if err := c.LoadConfig("/etc/autossh/autossh.yaml"); err != nil {
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
		if endpoint.Port == "" {
			endpoint.Port = c.Global.Port
		}
		if endpoint.User == "" {
			endpoint.User = c.Global.User
		}
		if endpoint.Password == "" {
			endpoint.Password = c.Global.Password
		}
		if endpoint.Passwords == nil {
			endpoint.Passwords = c.Global.Passwords
		}
		if endpoint.Key == "" {
			endpoint.Key = c.Global.Key
		}
		if endpoint.QAs == nil {
			endpoint.QAs = c.Global.QAs
		}
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
func (c *Config) ConsoleAdd() {
	input := ""
	endpoint := ssh.NewEndpoint()
	fmt.Printf("输入主机名称[" + c.Global.Name + "]: ")
	fmt.Scanln(&input)
	endpoint.Name = input

	log.ConsoleWithGreen("输入主机地址: ")
	fmt.Scanln(&input)
	endpoint.Host = input

	log.ConsoleWithGreen("输入IP地址: ")
	fmt.Scanln(&input)
	endpoint.IP = input

	log.ConsoleWithGreen("输入端口: ")
	fmt.Scanln(&input)
	endpoint.Port = input

	log.ConsoleWithGreen("输入用户名: ")
	fmt.Scanln(&input)
	endpoint.User = input

	log.ConsoleWithGreen("输入密码: ")
	fmt.Scanln(&input)
	endpoint.Password = input

	log.ConsoleWithGreen("输入证书: ")
	fmt.Scanln(&input)
	endpoint.Key = input

	c.Servers = append(c.Servers, endpoint)
	c.result = append(c.result, endpoint)

}
