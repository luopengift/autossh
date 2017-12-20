package core

import (
	"fmt"
	"github.com/luopengift/autossh/version"
	"github.com/luopengift/golibs/logger"
	"strings"
	"time"
)

func StartConsole(serverList *ServerList) error {
	//stdin := bufio.NewReader(os.Stdin)
	var input string
	for {
		logger.Info("Autossh... %s", time.Now().Format("2006/01/02 15:04:05"))
		serverList.Println()
		//os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		fmt.Printf("输入需要登录的服务器: ")
		//var input string
		n, err := fmt.Scanln(&input)
		if err != nil {
			logger.Error("input error: %v, %v", n, err)
			continue
		}
		//fmt.Println("%v,%v", input, n)
		//n, err := fmt.Fscanf(os.Stdin, "%s", &input)
		//inputReader := bufio.NewReader(os.Stdin)
		//input, err := inputReader.ReadString('\n')
		//input = strings.Trim(input, "\n")
		//fmt.Fscan(stdin, &input)
		//stdin.ReadString('\n')

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
			//fmt.Println("q|quit|exit: 退出")
			//fmt.Println("dump: 存储配置文件")
			//fmt.Println("add: 新增一台主机")
			//fmt.Println("rm: 删除一台主机")
			//fmt.Println()
		default:
			if strings.HasPrefix(input, "/") {
				result := serverList.Search(strings.Trim(input, "/"))
				switch len(result) {
				case 0:
					serverList.Reset()
				case 1:
					err = result[0].StartTerminal()
					serverList.Reset()
					return err
				}
			} else if false {
				continue
			} else {
				result := serverList.Match(strings.Trim(input, " "))
				switch len(result) {
				case 0:
					serverList.Reset()
				case 1:
					err = result[0].StartTerminal()
					serverList.Reset()
					return err
				}
			}
		}
		fmt.Println("end=", err)
	}
	return nil
}
