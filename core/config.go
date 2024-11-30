package core

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

type SSHConfig struct {
	Host string
}

var Config map[string]interface{}

func init() {
	_, err := toml.DecodeFile("pgscale.toml", &Config)
	cobra.CheckErr(err)
}

func panicNotFound(keys []string) {
	panic(fmt.Sprintf("Config for %v not found", keys))
}

func GetConfig(keys ...string) (interface{}, bool) {
	for i, key := range keys {
		val, found := Config[key]

		if !found {
			panicNotFound(keys)
		}

		if i == len(keys)-1 {
			return val, true
		}

		Config, found = val.(map[string]interface{})

		if !found {
			panicNotFound(keys)
		}
	}

	panicNotFound(keys)
	return nil, false
}
