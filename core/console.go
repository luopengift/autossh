package core

import (
	"bufio"
	"fmt"
	"github.com/luopengift/autossh/version"
	"github.com/luopengift/golibs/ssh"
	"github.com/luopengift/log"
	"os"
	"strings"
	"time"
)

func getInput() (string, error) {
	fmt.Printf("输入需要登录的服务器: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input[:len(input)-1], nil
}

func StartConsole(serverList *ServerList) error {
	for {
		log.Warn("Autossh... %s", time.Now().Format("2006/01/02 15:04:05"))
		serverList.Println()
		input, err := getInput()
		if err != nil {
			log.Error("input error: %v, %v", input, err)
			continue
		}
		fmt.Println("searching...")
		switch input {
		case "":
			continue
		case "-v", "-version":
			fmt.Println(version.VERSION)
		case "add":
			serverList.ConsoleAdd()
		case "show":
			continue
		case "dump":
			continue
		case "q", "quit", "exit":
			fmt.Println("exit...")
			return nil
		case "h", "help":
			fmt.Println("help....")
			fmt.Println("输入序号/名称/IP地址均可")
			fmt.Println("以'/'开头表示查询")
			fmt.Println("q|quit|exit: 退出")
			//fmt.Println("dump: 存储配置文件")
			//fmt.Println("add: 新增一台主机")
			//fmt.Println("rm: 删除一台主机")
			fmt.Println("\n")
		default:
			log.Warn("查询[%s]中,请稍后...", input)
			var result []*ssh.Endpoint
			switch input[0] {
			case '/':
				result = serverList.Search(strings.TrimSpace(string(input[1:])))
			default:
				result = serverList.Match(strings.TrimSpace(input))
			}

			switch len(result) {
			case 1:
				fmt.Println("正在登录", result[0].Ip)
				err = result[0].StartTerminal()
				serverList.Reset()
				return err
			default:
				serverList.Reset()
			}
		}
		//fmt.Println("end=", err)
	}
	return nil
}
