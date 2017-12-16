package core

import (
	"fmt"
	"github.com/luopengift/types"
	"os"
	//	"syscall"
	"bufio"
	"strings"
	"time"
)

func Exec() error {

	serverList := &ServerList{}
	err := types.ParseConfigFile("autossh.yml", serverList)
	if err != nil {
		return err
	}
	serverList.UseGlobalValues()
	serverList.Reset()
	stdin := bufio.NewReader(os.Stdin)
	var input string
	for {
		fmt.Println("Autossh...", time.Now().Format("2006/01/02 15:04:05"))
		serverList.Println()
		//os.Stdin = os.NewFile(uintptr(syscall.Stdin), "/dev/stdin")
		fmt.Printf("输入需要登录的服务器: ")
		//var input string
		//n, err := fmt.Scanf("%s\n", &input)
		//n, err := fmt.Fscanf(os.Stdin, "%s", &input)
		//inputReader := bufio.NewReader(os.Stdin)
		//input, err := inputReader.ReadString('\n')
		//input = strings.Trim(input, "\n")
		fmt.Fscan(stdin, &input)
		stdin.ReadString('\n')

		if err != nil {
			fmt.Println(">>", input, err)
		}
		fmt.Println("==|" + input + "|==")
		switch input {
		case "":
			continue
		case "-v", "-version":
			fmt.Println("v0.0.1_121617_beta")
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
