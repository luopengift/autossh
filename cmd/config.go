package cmd

import (
	"github.com/luopengift/autossh/config"
	"github.com/luopengift/types"
)

// LoadConfig config
func LoadConfig() (*config.Config, error) {
	conf := &config.Config{}
	if err := types.ParseConfigFile("/etc/autossh/config.yaml", conf); err != nil {
		return nil, err
	}
	return conf, nil
}
