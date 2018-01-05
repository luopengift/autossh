package main

import (
	"github.com/luopengift/golibs/logger"
	"github.com/luopengift/autossh/cmd"
)

func main() {
	logger.SetTimeFormat("")
	logger.SetLevel(logger.NULL)
	if err := cmd.Exec(); err != nil {
		logger.Error("%v", err)
	}
}
