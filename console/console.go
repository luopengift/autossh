package console

import (
	"context"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/autossh/pkg/endpoint"
	"github.com/luopengift/autossh/pkg/runtime"
	"github.com/luopengift/log"
	"github.com/luopengift/ssh"
	"github.com/luopengift/version"
)

// Welcome first time into console
func welcome() {
	log.ConsoleWithBlue("### 欢迎使用Autossh Jump System[%s] ###", time.Now().Format("2006/01/02 15:04:05"))
	log.ConsoleWithGreen("")
	for idx, v := range []string{
		"输入 P/p 查看主机列表.",
		"输入 G/g 查看主机分组",
		"输入 s + IP 直接登录.",
		"输入 V/v 查看版本号.",
		"输入 H/h 帮助.",
		"输入 Q/q 退出.",
	} {
		log.ConsoleWithGreen("\t%d) %s", idx, v)
	}
	log.ConsoleWithGreen("")
}

func searchEndpoints(ins *readline.Instance, endpoints endpoint.Endpoints) (endpoint.Endpoints, error) {
	endpoints.Print()
	if len(endpoints) == 1 {
		return endpoints, nil
	}
	ins.SetPrompt(readline.StaticPrompt("输入ID/IP/主机> "))
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	if isExit(input) {
		return nil, log.Errorf("exit")
	}
	result := endpoints.Search(input)
	switch len(result) {
	case 0:
		return searchEndpoints(ins, endpoints)
	case 1:
		return result, nil
	default:
		return searchEndpoints(ins, result)
	}

}

func searchGroups(ins *readline.Instance, groups *endpoint.Groups) (*endpoint.Groups, error) {
	groups.Print()
	if len(groups.List) == 1 {
		return groups, nil
	}
	ins.SetPrompt(readline.StaticPrompt("输入ID/主机组> "))
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	if isExit(input) {
		return nil, log.Errorf("exit")
	}
	result := groups.Search(input)
	switch len(result.List) {
	case 0:
		return searchGroups(ins, groups)
	case 1:
		return result, nil
	default:
		return searchGroups(ins, result)
	}
}

// Users users
type Users []string

// Print print users
func (s Users) Print() {
	log.ConsoleWithGreen("ID\t用户名")
	for idx, user := range s {
		log.ConsoleWithGreen("[%v]\t%s", idx, user)
	}
}

// Search search
func (s Users) Search(input string) Users {
	var result Users
	for index, user := range s {
		if input == strconv.Itoa(index) || strings.Contains(user, input) {
			result = append(result, user)
		}
	}
	return result
}

func searchUsers(ins *readline.Instance, users Users) ([]string, error) {
	if len(users) == 1 {
		return users, nil
	}
	users.Print()
	ins.SetPrompt(readline.StaticPrompt("输入用户/用户序号> "))
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	if isExit(input) {
		return nil, log.Errorf("exit")
	}
	// search
	result := users.Search(input)
	switch len(result) {
	case 0:
		return searchUsers(ins, users)
	case 1:
		return result, nil
	default:
		return searchUsers(ins, result)
	}
}

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
	welcome()
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
			welcome()
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
