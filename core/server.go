package core

import (
	"fmt"
	"github.com/luopengift/golibs/ssh"
	"github.com/luopengift/log"
	"strconv"
	"strings"
)

type ServerList struct {
	Global  *ssh.Endpoint   `yaml:"global"`
	Servers []*ssh.Endpoint `yaml:"servers"`
	result  []*ssh.Endpoint
}

func (s *ServerList) UseGlobalValues() {
	for _, endpoint := range s.Servers {
		if endpoint.Port == 0 {
			endpoint.Port = s.Global.Port
		}
		if endpoint.User == "" {
			endpoint.User = s.Global.User
		}
		if endpoint.Password == "" {
			endpoint.Password = s.Global.Password
		}
		if endpoint.Key == "" {
			endpoint.Key = s.Global.Key
		}
	}
}

func (s *ServerList) Println() {
	log.ConsoleWithGreen(fmt.Sprintf("%-4s\t%-20s\t%-40s\t%-5s", "序号", "名称", "IP", "端口"))
	for index, endpoint := range s.result {
		item := fmt.Sprintf("%-4d\t%-20s\t%-40s\t%-5d", index, endpoint.Name, endpoint.Ip, endpoint.Port)
		log.ConsoleWithGreen(item)
	}
}

func (s *ServerList) Reset() []*ssh.Endpoint {
	s.result = s.Servers
	return s.result
}

func (s *ServerList) Match(match string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for index, endpoint := range s.result {
		if match == strconv.Itoa(index) || match == endpoint.Name || match == endpoint.Host || match == endpoint.Ip {
			result = append(result, endpoint)
		}
	}
	s.result = result
	return s.result
}

func (s *ServerList) Search(search string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for _, endpoint := range s.result {
		if strings.Contains(endpoint.Name, search) || strings.Contains(endpoint.Host, search) || strings.Contains(endpoint.Ip, search) {
			result = append(result, endpoint)
		}
	}
	s.result = result
	return s.result
}

func (s *ServerList) Add(name, host, ip string, port int, user, password, key string) error {
	endpoint := ssh.NewEndpointWithValue(name, host, ip, port, user, password, key)
	s.Servers = append(s.Servers, endpoint)
	return nil
}

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
	endpoint.Ip = input

	log.ConsoleWithGreen("输入端口: ")
	fmt.Scanln(&input)
	endpoint.Port = 22

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
