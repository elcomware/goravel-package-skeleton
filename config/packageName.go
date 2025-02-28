package config

import (
	"github.com/goravel/framework/facades"
)

// config for VendorName/PackageName
func init() {
	config := facades.Config()
	config.Add("package_slug", map[string]any{
		"name": "package_slug",
	})
}
