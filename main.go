package main

import (
	"github.com/luopengift/autossh/cmd"
	"github.com/luopengift/log"
)

func main() {
	log.GetLogger("__ROOT__").SetFormatter(log.NewTextFormat("MESSAGE", 1))
	if err := cmd.Exec(); err != nil {
		log.Error("%v", err)
	}
}
