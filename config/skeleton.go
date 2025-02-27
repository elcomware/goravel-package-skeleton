package config

import (
	"github.com/goravel/framework/facades"
)

// config for VendorName/Skeleton
func init() {
	config := facades.Config()
	config.Add(":package_slug", map[string]any{
		"name": ":package_name",
	})
}
