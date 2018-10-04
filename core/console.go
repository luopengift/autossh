package core

import (
	"context"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

// Welcome first time into console
func welcome() {
	log.ConsoleWithBlue("\t\t欢迎使用Autossh Jump System")
	log.ConsoleWithGreen("")
}

// StartConsole StartConsole
func StartConsole(serverList *ServerList) error {
	log.Warn("Autossh... %s", time.Now().Format("2006/01/02 15:04:05"))
	welcome()
	rl, err := readline.New("> ")
	if err != nil {
		return err
	}
	defer rl.Close()
	for {
		input, err := rl.Readline()
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		switch {
		case input == "list":
			serverList.println()
		case input == "v", input == "version", input == "-v", input == "-version", input == "--version":
			log.ConsoleWithGreen("version:%v, build time:%v, build tag:%v", version.VERSION, version.TIME, version.GIT)
		case input == "add":
			serverList.ConsoleAdd()
		case input == "show":
			continue
		case input == "dump":
			continue
		case input == "q", input == "quit", input == "exit":
			log.ConsoleWithGreen("exit...")
			return nil
		case input == "h", input == "help":
			log.ConsoleWithGreen("help....")
			log.ConsoleWithGreen("输入序号/名称/IP地址均可")
			log.ConsoleWithGreen("以'/'开头表示查询")
			log.ConsoleWithGreen("q|quit|exit: 退出")
			//log.ConsoleWithGreen("dump: 存储配置文件")
			//log.ConsoleWithGreen("add: 新增一台主机")
			//log.ConsoleWithGreen("rm: 删除一台主机")
			log.ConsoleWithGreen("\n")
		case input == "qa":
			// question and answer
		case strings.HasPrefix(input, "search"), strings.HasPrefix(input, "s "):
			inputList := strings.Split(input, " ")
			if len(inputList) != 2 {
				log.ConsoleWithGreen("not 2...")
				serverList.println()
				continue
			}

			log.ConsoleWithGreen("查询[%s]中,请稍后...", inputList[1])
			var result []*ssh.Endpoint
			result = serverList.Match(inputList[1])
			if len(result) == 1 {
				log.ConsoleWithGreen("正在登录 %v", result[0].IP)
				err = result[0].StartTerminal()
			}
			serverList.Reset()

		default:
			ctx := context.TODO()
			Bash(ctx, input, nil)
		}
		//log.ConsoleWithGreen("end=%v", err)
	}
}
