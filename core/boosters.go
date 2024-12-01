package core

import "fmt"

func panicNotFound(keys []string) {
	panic(fmt.Sprintf("Config for %v not found", keys))
}

func AsStrMap(v interface{}) map[string]interface{} {
	return v.(map[string]interface{})
}
