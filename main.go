package main

import (
	"os"
	"fmt"
	"github.com/luopengift/types"
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/golibs/ssh"
	"strconv"
	"strings"
)

var log *logger.Logger

type Config struct {
    Servers []*ssh.Endpoint `yaml:"servers"`
}

func main() {
	logger.SetPrefix("")
	logger.Info("Autossh...")
	err := Exec()
	fmt.Println(err)
}
func Exec() error {

    fd, err := os.OpenFile("log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    log = logger.NewLogger("2006/01/02 15:04:05.000", logger.DEBUG,fd)
    conf := &Config{}
    err = types.ParseConfigFile("autossh.yml", conf)
    if err != nil {
        return err
    }

	endpoints := conf.Servers
	input := "" //屏幕输入
	for {
		Println(endpoints)
		fmt.Printf("输入需要登录的服务器: ")
		fmt.Scanln(&input)
		switch input {
		case "q", "quit", "exit":
			fmt.Println("exit...")
			return nil
		case "h", "help":
			fmt.Println("help")
		default:
			if strings.HasPrefix(input, "/") {
				result := Search(endpoints, strings.Trim(input, "/"))
				switch len(result) {
				case 0:
					result = endpoints
				case 1:
					return result[0].StartTerminal()
				}
				endpoints = result
			} else if false {
				continue
			} else {
				result := Match(endpoints, strings.Trim(input, " "))
				switch len(result) {
				case 0:
					result = endpoints
				case 1:
					return result[0].StartTerminal()
				}
				endpoints = result
			}
		}
	}
	return nil
}

func Println(servers []*ssh.Endpoint) {
	fmt.Println(fmt.Sprintf("序号\t名称\tIP\t\t端口"))
	for index, endpoint := range servers {
		item := fmt.Sprintf("%v\t%v\t%v\t\t%v", index, endpoint.Name, endpoint.Ip, endpoint.Port)
		fmt.Println(item)
	}
}

func Match(servers []*ssh.Endpoint, match string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for index, endpoint := range servers {
		if match == strconv.Itoa(index) || match == endpoint.Host || match == endpoint.Ip {
			result = append(result, endpoint)
		}
	}
	return result

}

func Search(servers []*ssh.Endpoint, search string) []*ssh.Endpoint {
	result := []*ssh.Endpoint{}
	for _, endpoint := range servers {
		if strings.Contains(endpoint.Host, search) || strings.Contains(endpoint.Ip, search) {
			result = append(result, endpoint)
		}
	}
	return result
}
