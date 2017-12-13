package main

import (
	"fmt"
	"github.com/luopengift/autossh/core"
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/types"
	"os"
	"strings"
	"time"
)

var log *logger.Logger

func main() {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return
	}

	log = logger.NewLogger("", logger.DEBUG, f)
	err = Exec()
	if err != nil {
		fmt.Println(err)
	}
}
func Exec() error {

	serverList := &core.ServerList{}
	err := types.ParseConfigFile("autossh.yml", serverList)
	if err != nil {
		return err
	}

	serverList.Reset()
	for {
		fmt.Println("Autossh...", time.Now().Format("2006/01/02 15:04:05"))
		serverList.Println()
		fmt.Printf("输入需要登录的服务器: ")
		var input string
		n, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(n, err)
		}
		switch input {
		case "":
			continue
		case "q", "quit", "exit":
			fmt.Println("exit...")
			return nil
		case "h", "help":
			fmt.Println("help")
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

