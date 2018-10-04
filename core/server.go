package core

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
)

// ServerList serverList
type ServerList struct {
	Global  *ssh.Endpoint   `yaml:"global"`
	Servers []*ssh.Endpoint `yaml:"servers"`
	result  []*ssh.Endpoint
}

// UseGlobalValues UseGlobalValues
func (s *ServerList) UseGlobalValues() {
	for _, endpoint := range s.Servers {
		if endpoint.Port == "" {
			endpoint.Port = s.Global.Port
		}
		if endpoint.User == "" {
			endpoint.User = s.Global.User
		}
		if endpoint.Password == "" {
			endpoint.Password = s.Global.Password
		}
		if endpoint.Passwords == nil {
			endpoint.Passwords = s.Global.Passwords
		}
		if endpoint.Key == "" {
			endpoint.Key = s.Global.Key
		}
		if endpoint.QAs == nil {
			endpoint.QAs = s.Global.QAs
		}
	}
}

//
func (s *ServerList) println() {
	format := "%-4v\t%-20s\t%-40s\t%-5s"
	log.ConsoleWithGreen(fmt.Sprintf(format, "序号", "名称", "地址", "用户名"))
	for index, endpoint := range s.result {
		log.ConsoleWithGreen(
			fmt.Sprintf(format, index, endpoint.Name, endpoint.Address(), endpoint.User),
		)
	}
}

// Reset Reset
func (s *ServerList) Reset() []*ssh.Endpoint {
	s.result = s.Servers
	return s.result
}

// Match match
func (s *ServerList) Match(match string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for index, endpoint := range s.result {
		if match == strconv.Itoa(index) || match == endpoint.Name || match == endpoint.Host || match == endpoint.IP {
			result = append(result, endpoint)
		}
	}
	s.result = result
	return s.result
}

// Search search
func (s *ServerList) Search(search string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for index, endpoint := range s.result {
		if search == strconv.Itoa(index) || strings.Contains(endpoint.Name, search) || strings.Contains(endpoint.Host, search) || strings.Contains(endpoint.IP, search) {
			result = append(result, endpoint)
		}
	}
	s.result = result
	return s.result
}

// Add add
func (s *ServerList) Add(name, host, ip, port, user, password, key string) error {
	endpoint := ssh.NewEndpointWithValue(name, host, ip, port, user, password, key)
	s.Servers = append(s.Servers, endpoint)
	return nil
}

// ConsoleAdd ConsoleAdd
func (s *ServerList) ConsoleAdd() {
	input := ""
	endpoint := ssh.NewEndpoint()
	fmt.Printf("输入主机名称[" + s.Global.Name + "]: ")
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

	s.Servers = append(s.Servers, endpoint)
	s.result = append(s.result, endpoint)

}
