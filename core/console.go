package core

import (
	"context"
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
	log.ConsoleWithGreen("\t1) 输入ID直接登陆.")
	log.ConsoleWithGreen("\t2) 输入P/p查看机器列表.")
	log.ConsoleWithGreen("\t2) 输入s + IP, 主机名搜索.")
	log.ConsoleWithGreen("\t3) 输入H/h帮助.")
	log.ConsoleWithGreen("\t4) 输入Q/q退出.")
	log.ConsoleWithGreen("")
}

type Prompt struct {}
func (p *Prompt) String() string {
	return "> "
	//return time.Now().Format("2006/01/02 15:04:05") + "> "
}

// StartConsole StartConsole
func StartConsole(ctx context.Context, conf *config.Config) error {
	log.Warn("Autossh... %s", time.Now().Format("2006/01/02 15:04:05"))
	welcome()
	rl, err := readline.New(&Prompt{})
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
		case input == "v", input == "version", input == "-v", input == "-version", input == "--version":
			log.ConsoleWithGreen("version:%v, build time:%v, build tag:%v", version.VERSION, version.TIME, version.GIT)
		case input == "add":
			conf.ConsoleAdd()
		case input == "show":
			continue
		case input == "dump":
			continue
		case input == "q", input == "Q", input == "quit", input == "exit":
			log.ConsoleWithGreen("exit...")
			return nil
		case input == "h", input == "H", input == "help":
			welcome()
			// log.ConsoleWithGreen("help....")
			// log.ConsoleWithGreen("输入序号/名称/IP地址均可")
			// log.ConsoleWithGreen("以'/'开头表示查询")
			// log.ConsoleWithGreen("q|quit|exit: 退出")
			// //log.ConsoleWithGreen("dump: 存储配置文件")
			// //log.ConsoleWithGreen("add: 新增一台主机")
			// //log.ConsoleWithGreen("rm: 删除一台主机")
			// log.ConsoleWithGreen("\n")
		case input == "qa":
			// question and answer
		case strings.HasPrefix(input, "s "):
			inputList := strings.Split(input, " ")
			if len(inputList) != 2 {
				log.ConsoleWithGreen("not 2...")
				//conf.Println()
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
