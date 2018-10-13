package core

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

// Welcome first time into console
func welcome() {
	log.ConsoleWithBlue("\t\t### 欢迎使用Autossh Jump System ###")
	log.ConsoleWithGreen("")
	for idx, v := range []string{
		"输入P/p 查看机器列表.",
		"输入s + IP,主机名 搜索.",
		"输入V/v 查看版本号.",
		"输入H/h 帮助.",
		"输入Q/q 退出.",
	} {
		log.ConsoleWithGreen("\t%d) %s", idx, v)
	}
	log.ConsoleWithGreen("")
}

// StartConsole StartConsole
func StartConsole(ctx context.Context, conf *config.Config) error {
	log.Warn("Autossh... %s", time.Now().Format("2006/01/02 15:04:05"))
	welcome()
	rl, err := readline.New(readline.StaticPrompt(fmt.Sprintf("%s> ", os.Getenv("USER"))))
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
		case input == "P", input == "p":
			conf.Println()
		case input == "V", input == "v", input == "version", input == "-v", input == "-version", input == "--version":
			log.ConsoleWithGreen("version: %v, buildTime: %v, buildTag: %v", version.VERSION, version.TIME, version.GIT)
		case input == "add": // 新增一台主机
			conf.ConsoleAdd()
		case input == "show":
			continue
		case input == "dump": // 存储配置文件
			continue
		case input == "q", input == "Q", input == "quit", input == "exit":
			log.ConsoleWithGreen("exit...")
			return nil
		case input == "h", input == "H", input == "help":
			welcome()
		case input == "qa":
			// question and answer
		case strings.HasPrefix(input, "s "):
			inputList := strings.Split(input, " ")
			if len(inputList) != 2 {
				log.ConsoleWithRed("输入有误! 请输入H/h 查看帮助.")
				continue
			}
			log.ConsoleWithGreen("查询[%s]中,请稍后...", inputList[1])
			var result []*ssh.Endpoint
			result = conf.Match(inputList[1])
			if len(result) == 1 {
				log.ConsoleWithGreen("正在登录 %v", result[0].IP)
				if conf.Backup != "" {
					filepath := path.Join(conf.Backup, os.Getenv("USER"), result[0].IP+time.Now().Format("20060102150405.log"))
					f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
					if err != nil {
						log.Error("%v", err)
						continue
					}
					defer f.Close()
					result[0].SetWriters(f)
				}
				if err = result[0].StartTerminal(); err != nil {
					log.ConsoleWithRed("%v", err)
				}
			}
			conf.Reset()
		default:
			if conf.Shell {
				Bash(ctx, input, nil)
			}
		}
		//log.ConsoleWithGreen("end=%v", err)
	}
}
