package core

import (
	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var Config map[string]interface{}

func init() {
	_, err := toml.DecodeFile("pgscale.toml", &Config)
	cobra.CheckErr(err)
}

func GetConfig(keys ...string) map[string]interface{} {
	currentConfig := Config

	for i, key := range keys {
		val, ok := currentConfig[key]

		if !ok {
			panicNotFound(keys)
		}

		if i == len(keys)-1 {
			return AsStrMap(val)
		}

		currentConfig = AsStrMap(val)
	}

	return nil
}
