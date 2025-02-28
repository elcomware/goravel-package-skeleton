package config

import (
	"github.com/goravel/framework/facades"
)

// config for VendorName/PackageName
func init() {
	config := facades.Config()
	config.Add("packageName", map[string]any{
		"name": "packageName",
	})
}
