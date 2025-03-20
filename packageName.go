package packageName

import (
	"github.com/goravel/framework/facades"
	"github.com/vendorName/packageName/tools"
	"golang.org/x/term"
	"os"

	"path/filepath"
	"strings"
)

// PackageName represents a modular component in Goravel.
type PackageName struct {
	FullName                string
	ShortName               string
	BasePath                string
	PublishableProviderName string

	ConfigTools         tools.ConfigTools
	ViewComponentTools  tools.ViewComponentTools
	ViewComposerTools   tools.ViewComposerTools
	ViewSharedDataTools tools.ViewSharedDataTools
	MigrationTools      tools.MigrationTools
	CommandTools        tools.CommandTools
	AssetTools          tools.AssetTools
	TranslationTools    tools.TranslationTools
	ViewsTools          tools.ViewTools
	RouteTools          tools.RoutesTools
	ProviderTools       tools.ProviderTools
}

// NewPackageName creates a new PackageName instance.
func NewPackageName(name string) *PackageName {

	shortName := getShortName(name)

	return &PackageName{
		BasePath:  "",
		FullName:  name,
		ShortName: shortName,

		// Initialize dependencies if needed
		//Assets:          &tools.Assets{},
		ConfigTools:         tools.ConfigTools{ShortName: shortName},
		ViewComponentTools:  tools.ViewComponentTools{Components: make(map[string]string)},
		ViewComposerTools:   tools.ViewComposerTools{Composers: make(map[string]string)},
		ViewSharedDataTools: tools.ViewSharedDataTools{SharedData: make(map[string]interface{})},
	}
}

// getShortName is a function to dynamically determine the ShortName
// ShortName returns the short FullName of the package, removing the 'goravel-' prefix.
func getShortName(name string) string {
	// Example logic to determine the short FullName
	// This can fetch it from the config, environment, or fallback to a default value
	if facades.Config().GetString("package.short_name") != "" {
		return facades.Config().GetString("package.short_name")
	}

	if name == "" {
		name = "default"
	}
	return strings.TrimPrefix(name, "goravel-")
}

// GetBasePath returns the base path of the package, optionally appending a directory.
func (p *PackageName) GetBasePath(directory ...string) string {
	if len(directory) == 0 {
		return p.BasePath
	}
	return filepath.Join(append([]string{p.BasePath}, directory...)...)
}

// SetBasePath sets the base path of the package.
func (p *PackageName) SetBasePath(path string) *PackageName {
	p.BasePath = path
	return p
}

// IsRunningInConsole checks if the application is running in a terminal/console.
func (p *PackageName) IsRunningInConsole() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}
