package main

import (
	"context"

	"github.com/luopengift/autossh/cmd"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/log"
)

func main() {
	log.GetLogger("__ROOT__").SetFormatter(log.NewTextFormat("MESSAGE", log.ModeColor))
	conf := config.Init()
	conf.LoadRootConfig()
	if err := cmd.Exec(context.Background(), conf); err != nil {
		log.Error("%v", err)
	}
}
