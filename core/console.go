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
	fmt.Fprintf(os.Stderr, "输入需要登录的服务器: ")
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
		log.ConsoleWithGreen("searching...")
		switch input {
		case "":
			continue
		case "-v", "-version":
			log.ConsoleWithGreen(version.VERSION)
		case "add":
			serverList.ConsoleAdd()
		case "show":
			continue
		case "dump":
			continue
		case "q", "quit", "exit":
			log.ConsoleWithGreen("exit...")
			return nil
		case "h", "help":
			log.ConsoleWithGreen("help....")
			log.ConsoleWithGreen("输入序号/名称/IP地址均可")
			log.ConsoleWithGreen("以'/'开头表示查询")
			log.ConsoleWithGreen("q|quit|exit: 退出")
			//log.ConsoleWithGreen("dump: 存储配置文件")
			//log.ConsoleWithGreen("add: 新增一台主机")
			//log.ConsoleWithGreen("rm: 删除一台主机")
			log.ConsoleWithGreen("\n")
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
				log.ConsoleWithGreen("正在登录 %v", result[0].Ip)
				err = result[0].StartTerminal()
				serverList.Reset()
				return err
			default:
				serverList.Reset()
			}
		}
		//log.ConsoleWithGreen("end=%v", err)
	}
	return nil
}
