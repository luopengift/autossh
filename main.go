package main

import (
	"fmt"
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/golibs/ssh"
	"github.com/luopengift/types"
	"strconv"
	"strings"
)

var log *logger.Logger

type ServerList struct {
	Global  *ssh.Endpoint   `yaml:"global"`
	Servers []*ssh.Endpoint `yaml:"servers"`
	result  []*ssh.Endpoint
}

func (s *ServerList) Println() {
	fmt.Println(fmt.Sprintf("序号\t名称\tIP\t\t端口"))
	for index, endpoint := range s.result {
		item := fmt.Sprintf("%v\t%v\t%v\t\t%v", index, endpoint.Name, endpoint.Ip, endpoint.Port)
		fmt.Println(item)
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

func recv_console() string {
	input := ""
	n, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("scan error:", err)
	}
	fmt.Println("scan len", n)
	return input
}

func ScanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}
	return string(b)
}

func main() {
	logger.SetPrefix("")
	logger.Info("Autossh...")
	err := Exec()
	if err != nil {
		fmt.Println(err)
	}
}
func Exec() error {

	serverList := &ServerList{}
	err := types.ParseConfigFile("autossh.yml", serverList)
	if err != nil {
		return err
	}

	serverList.Reset()
	input := ""
	for {

		serverList.Println()
		fmt.Println("")
		fmt.Printf("输入需要登录的服务器: ")
		input = recv_console()
		fmt.Println("input: ", "|"+input+"|")

		switch input {
		case "q", "quit", "exit":
			fmt.Println("exit...")
			return nil
		case "h", "help":
			fmt.Println("help")
		default:
			fmt.Println("default")
			if strings.HasPrefix(input, "/") {
				result := serverList.Search(strings.Trim(input, "/"))
				switch len(result) {
				case 0:
					serverList.Reset()
				case 1:
					err = result[0].StartTerminal()
					fmt.Println("close.", err)
					serverList.Reset()
				}
			} else if false {
				continue
			} else {
				result := serverList.Match(strings.Trim(input, " "))
				switch len(result) {
				case 0:
					serverList.Reset()
				case 1:
					result[0].StartTerminal()
					fmt.Println("close.", err)
					serverList.Reset()
				}
			}
		}
	}
	return nil
}
