package core

import (
	"github.com/luopengift/golibs/ssh"
	"fmt"
	"strconv"
	"strings"
)


type ServerList struct {
	Global  *ssh.Endpoint   `yaml:"global"`
	Servers []*ssh.Endpoint `yaml:"servers"`
	result  []*ssh.Endpoint
}

func (s *ServerList) Println() {
	fmt.Println("==================================")
	fmt.Println(fmt.Sprintf("序号\t名称\t主机\tIP\t端口"))
	for index, endpoint := range s.result {
		item := fmt.Sprintf("%v\t%v\t%v\t%v\t%v", index, endpoint.Name, endpoint.Host, endpoint.Ip, endpoint.Port)
		fmt.Println(item)
	}
	fmt.Println("==================================")
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
	fmt.Printf("输入主机名称[%v]: ", s.Global)
	fmt.Scanln(&input)
	s.Global.Name = input
	fmt.Println(endpoint)

}


