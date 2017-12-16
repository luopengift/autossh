package main

import (
	"fmt"
	"github.com/luopengift/autossh/core"
	"github.com/luopengift/golibs/logger"
	"os"
)

var log *logger.Logger

func main() {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		return
	}
	//ToFile()

	log = logger.NewLogger("", logger.DEBUG, f)
	err = core.Exec()
	if err != nil {
		fmt.Println(err)
	}

}
