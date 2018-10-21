package main

import (
	"context"

	"github.com/luopengift/autossh/cmd"
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/log"
)

func main() {
	conf := config.Init()
	conf.LoadRootConfig()
	log.SetLevel(log.INFO).SetTextFormat("MESSAGE", log.ModeColor)
	if err := cmd.Run(context.Background(), conf); err != nil {
		log.Error("%v", err)
	}
}
