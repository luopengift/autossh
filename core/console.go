package core

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

func getInput(ps string) (string, error) {
	fmt.Fprintf(os.Stderr, ps)
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input[:len(input)-1], nil
}

// Welcome first time into console
func welcome() {
	log.ConsoleWithBlue("\t\t欢迎使用Autossh Jump System")
	log.ConsoleWithGreen("")
}

// StartConsole StartConsole
func StartConsole(serverList *ServerList) error {
	log.Warn("Autossh... %s", time.Now().Format("2006/01/02 15:04:05"))
	welcome()
	PS := "> "
	for {
		//serverList.println()
		input, err := getInput(PS)
		if err != nil {
			log.Error("input error: %v, %v", input, err)
			continue
		}
		switch input {
		case "":
			continue
		case "list":
			serverList.println()
		case "v", "version", "-v", "-version", "--version":
			log.ConsoleWithGreen("version:%v, build time:%v, build tag:%v", version.VERSION, version.TIME, version.GIT)
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
		case "search":
			log.ConsoleWithGreen("searching...")
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
				log.ConsoleWithGreen("正在登录 %v", result[0].IP)
				err = result[0].StartTerminal()
				serverList.Reset()
				return err
			default:
				serverList.Reset()
			}
		default:
			ctx := context.TODO()
			Bash(ctx, input, nil)
		}
		//log.ConsoleWithGreen("end=%v", err)
	}
}
