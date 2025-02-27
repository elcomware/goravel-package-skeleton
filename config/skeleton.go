package config

import (
	"github.com/goravel/framework/facades"
)

// config for Elcomware/GoSecure
func init() {
	config := facades.Config()
	config.Add("go-secure", map[string]any{
		"name": "go-secure",
	})
}
