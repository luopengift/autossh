package console

import (
	"context"
	"os"
	"path"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/autossh/pkg/runtime"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

// Login login
func Login(endpoint *ssh.Endpoint, conf *config.Config) error {
	log.ConsoleWithGreen("正在使用用户[%s]登录 %s:%s", endpoint.User, endpoint.IP, endpoint.Port)
	if conf.Backup != "" {
		filepath := path.Join(conf.Backup, os.Getenv("USER"), endpoint.IP+time.Now().Format("20060102150405.log"))
		f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Error("%v", err)
			return err
		}
		defer f.Close()
		endpoint.SetWriters(f)
	}
	defer endpoint.Close()
	if err := endpoint.StartTerminal(); err != nil {
		log.ConsoleWithRed("%v", err)
		return err
	}
	return nil
}

// StartConsole StartConsole
func StartConsole(ctx context.Context, conf *config.Config) error {
	r, err := runtime.NewRuntime()
	if err != nil {
		return err
	}
	r.SetEndpoints(conf.Endpoints)
	help()
	rl, err := readline.New(r)
	if err != nil {
		return err
	}
	defer rl.Close()
	for {
		rl.SetPrompt(r)
		input, err := rl.Readline()
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		switch {
		case input == "P", input == "p":
			endpoints, err := searchEndpoints(rl, r.Endpoints)
			if err != nil {
				log.ConsoleWithRed("%v", err)
				continue
			}
			endpoint := endpoints[0].Copy()
			users, _ := endpoint.GetUsers()
			users, err = searchUsers(rl, users)
			if err != nil {
				log.ConsoleWithRed("%v", err)
				continue
			}
			endpoint.User = users[0]
			Login(endpoint, conf)
		case input == "G", input == "g":
			groups, err := searchGroups(rl, r.Groups)
			if err != nil {
				log.ConsoleWithRed("%v", err)
				continue
			}
			endpoints := groups.Groups[groups.List[0]]
			endpoints, err = searchEndpoints(rl, endpoints)
			if err != nil {
				log.ConsoleWithRed("%v", err)
				continue
			}
			endpoint := endpoints[0].Copy()
			users, _ := endpoint.GetUsers()
			users, err = searchUsers(rl, users)
			if err != nil {
				log.ConsoleWithRed("%v", err)
				continue
			}
			endpoint.User = users[0]
			Login(endpoint, conf)
		case isVersion(input):
			log.ConsoleWithGreen("version: %v, buildTime: %v, buildTag: %v", version.VERSION, version.TIME, version.GIT)
		case input == "add": // 新增一台主机
			if r.Super {
				if err = conf.ConsoleAdd(); err != nil {
					log.ConsoleWithRed("%v", err)
				}
				r.SetEndpoints(conf.Endpoints)
			}
		case input == "set-super":
			r.Super = true
			log.ConsoleWithMagenta("进入Super模式!")
		case input == "print config":
			log.Display("config", conf)
		case input == "print runtime":
			log.Display("runtime", r)
		case strings.HasPrefix(input, "dump "): // 存储配置文件
			if r.Super {
				inputList := strings.Split(input, " ")
				if len(inputList) != 2 {
					log.ConsoleWithRed("输入有误! 请输入 H/h 查看帮助.")
					continue
				}
				if err = conf.Dump(inputList[1]); err != nil {
					log.ConsoleWithRed("Dump 出错: %v", err)
				} else {
					log.ConsoleWithGreen("Dump 成功: %v", inputList[1])
				}
			}
		case isExit(input):
			if r.Super {
				r.Super = false
				log.ConsoleWithMagenta("退出Super模式.")
				continue
			}
			log.ConsoleWithGreen("exit...")
			return nil
		case isHelp(input):
			help()
		case input == "qa":
			// question and answer
		case strings.HasPrefix(input, "s "):
			inputList := strings.Split(input, " ")
			if len(inputList) != 2 {
				log.ConsoleWithRed("输入有误! 请输入 H/h 查看帮助.")
				continue
			}
			endpoint := getEndpoint(inputList[1])
			endpoint.Mask(conf.Global)
			Login(endpoint, conf)
		default:
			if conf.Shell {
				Bash(ctx, input, nil)
			}
		}
		//log.ConsoleWithGreen("end=%v", err)
	}
}
