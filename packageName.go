package packageName

import (
	"github.com/goravel/framework/contracts/config"
	"github.com/vendorName/packageName/tools"
	"golang.org/x/term"
	"os"

	"path/filepath"
	"strings"
)

// PackageName represents a modular component in Goravel.
type PackageName struct {
	config         config.Config
	name           string
	shortName      string
	basePath       string
	pkgConfig      tools.PackageConfigs
	viewComponents tools.PackageViewComponents
	viewComposer   tools.PackageViewComposers
	viewSharedData tools.PackageViewSharedData
	tools.PackageMigrations
	tools.PackageCommands
	tools.PackageAssets
	tools.PackageTranslations
	tools.PackageViews
	tools.PackageRoutes
	tools.PackageProviders
}

// NewPackageName creates a new PackageName instance.
func NewPackageName(config config.Config, name string) *PackageName {

	shortName := getShortNameFromConfig(config, name)

	return &PackageName{
		config:    config,
		basePath:  "",
		name:      name,
		shortName: shortName,

		// Initialize dependencies if needed
		//Assets:          &tools.Assets{},
		pkgConfig:      tools.PackageConfigs{ShortName: shortName},
		viewComponents: tools.PackageViewComponents{ViewComponents: make(map[string]string)},
		viewComposer:   tools.PackageViewComposers{ViewComposers: make(map[string]string)},
		viewSharedData: tools.PackageViewSharedData{SharedViewData: make(map[string]interface{})},
	}
}

// getShortNameFromConfig is a function to dynamically determine the ShortName
func getShortNameFromConfig(config config.Config, name string) string {
	// Example logic to determine the short name
	// This can fetch it from the config, environment, or fallback to a default value
	if config.GetString("package.short_name") != "" {
		return config.GetString("package.short_name")
	}

	if name == "" {
		name = "default"
	}
	return strings.TrimPrefix(name, "goravel-")
}

// ShortName returns the short name of the package, removing the 'goravel-' prefix.
func (p *PackageName) ShortName() string {
	if p.name == "" {
		p.name = "default"
	}
	return strings.TrimPrefix(p.name, "goravel-")
}

// GetBasePath returns the base path of the package, optionally appending a directory.
func (p *PackageName) GetBasePath(directory ...string) string {
	if len(directory) == 0 {
		return p.basePath
	}
	return filepath.Join(append([]string{p.basePath}, directory...)...)
}

// SetBasePath sets the base path of the package.
func (p *PackageName) SetBasePath(path string) *PackageName {
	p.basePath = path
	return p
}

// IsRunningInConsole checks if the application is running in a terminal/console.
func (p *PackageName) IsRunningInConsole() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}
