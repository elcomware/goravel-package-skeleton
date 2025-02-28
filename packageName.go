package packageName

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
)

type PackageName struct {
	config config.Config
}

func NewPackageName(config config.Config) *PackageName {
	return &PackageName{config: config}
}

func (s *PackageName) PackageAction() string {
	return fmt.Sprintf("Welcome To Goravel %s", s.config.GetString("package_slug.name"))
}
