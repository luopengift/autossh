package main

import (
	"github.com/luopengift/autossh/cmd"
	"github.com/luopengift/golibs/logger"
)

func main() {
	logger.SetTimeFormat("")
	logger.SetLevel(logger.NULL)
	if err := cmd.Exec(); err != nil {
		logger.Error("%v", err)
	}
}
