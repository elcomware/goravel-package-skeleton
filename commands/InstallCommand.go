package commands

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

// InstallCommand is a Goravel command for installing packages.
type InstallCommand struct {
	StartFunc                func(*InstallCommand) // Function to run at the start of installation
	Publishes                []string              // Tags for resources to publish (e.g., "config", "migrations")
	RunMigrations            bool                  // Whether to ask to run migrations
	CopyServiceProviderInApp bool                  // Whether to copy the service provider to the app
	StarRepo                 string                // GitHub repo to ask the user to star (e.g., "spatie/laravel-package-tools")
	EndFunc                  func(*InstallCommand) // Function to run at the end of installation
	Package                  SudoPackage
}

type SudoPackage struct {
	PublishableProviderName string
	ShortName               string
	FullName                string
}

// NewInstallCommand creates a new InstallCommand instance.
func NewInstallCommand(p SudoPackage) *InstallCommand {
	return &InstallCommand{
		Package: p,
	}
}

// Signature defines the command signature for Goravel.
func (ic *InstallCommand) Signature() string {
	return fmt.Sprintf("%s:install", ic.Package.ShortName)
}

// Description defines the command description for Goravel.
func (ic *InstallCommand) Description() string {
	return fmt.Sprintf("Install %s", ic.Package.FullName)
}

// Extend registers the command with Goravel.
func (ic *InstallCommand) Extend() command.Extend {
	return command.Extend{
		Flags: []command.Flag{}, // Add flags if needed
	}
}

// Handle executes the installation process.
func (ic *InstallCommand) Handle(ctx console.Context) error {
	// Run the start function if defined
	if ic.StartFunc != nil {
		ic.StartFunc(ic)
	}

	// Publish resources
	for _, tag := range ic.Publishes {
		name := strings.ReplaceAll(tag, "-", " ")
		ctx.Info(fmt.Sprintf("Publishing %s...", name))
		ic.callSilently("publish", "--tag", fmt.Sprintf("%s-%s", ic.Package.ShortName, tag))
	}

	// Ask to run migrations
	if ic.RunMigrations {
		if confirm(ctx, "Would you like to run the migrations now?") {
			ctx.Info("Running migrations...")
			ic.call("migrate")
		}
	}

	// Copy and register the service provider
	if ic.CopyServiceProviderInApp {
		ctx.Info("Publishing service provider...")
		ic.copyServiceProviderInApp()
	}

	// Ask to star the GitHub repo
	if ic.StarRepo != "" {
		if confirm(ctx, "Would you like to star our repo on GitHub?") {
			repoUrl := fmt.Sprintf("https://github.com/%s", ic.StarRepo)
			openBrowser(repoUrl)
		}
	}

	// Display success message
	ctx.Info(fmt.Sprintf("%s has been installed!", ic.Package.ShortName))

	// Run the end function if defined
	if ic.EndFunc != nil {
		ic.EndFunc(ic)
	}

	return nil
}

// Publish adds tags for resources to publish.
func (ic *InstallCommand) Publish(tags ...string) *InstallCommand {
	ic.Publishes = append(ic.Publishes, tags...)
	return ic
}

// PublishConfigFile publishes the configuration file.
func (ic *InstallCommand) PublishConfigFile() *InstallCommand {
	return ic.Publish("config")
}

// PublishAssets publishes asset files.
func (ic *InstallCommand) PublishAssets() *InstallCommand {
	return ic.Publish("assets")
}

// PublishMigrations publishes migration files.
func (ic *InstallCommand) PublishMigrations() *InstallCommand {
	return ic.Publish("migrations")
}

// RunMigrations enables the migration prompt.
func (ic *InstallCommand) AskToRunMigrations() *InstallCommand {
	ic.RunMigrations = true
	return ic
}

// CopyAndRegisterServiceProviderInApp enables copying the service provider.
func (ic *InstallCommand) CopyAndRegisterServiceProviderInApp() *InstallCommand {
	ic.CopyServiceProviderInApp = true
	return ic
}

// AskToStarRepoOnGitHub sets the GitHub repo to ask the user to star.
func (ic *InstallCommand) AskToStarRepoOnGitHub(vendorSlashRepoName string) *InstallCommand {
	ic.StarRepo = vendorSlashRepoName
	return ic
}

// StartWith sets the function to run at the start of installation.
func (ic *InstallCommand) StartWith(callable func(*InstallCommand)) *InstallCommand {
	ic.StartFunc = callable
	return ic
}

// EndWith sets the function to run at the end of installation.
func (ic *InstallCommand) EndWith(callable func(*InstallCommand)) *InstallCommand {
	ic.EndFunc = callable
	return ic
}

// copyServiceProviderInApp copies and registers the service provider.
func (ic *InstallCommand) copyServiceProviderInApp() *InstallCommand {
	providerName := ic.Package.PublishableProviderName
	if providerName == "" {
		return ic
	}

	// Simulate copying the service provider
	ic.callSilently("publish", "--tag", fmt.Sprintf("%s-provider", ic.Package.ShortName))

	// Simulate registering the provider
	fmt.Printf("Registered provider: %s\n", providerName)
	return ic
}

// call simulates running a console command.
func (ic *InstallCommand) call(command string, args ...string) {
	fmt.Printf("Calling command: %s %s\n", command, strings.Join(args, " "))
}

// callSilently simulates running a console command silently.
func (ic *InstallCommand) callSilently(command string, args ...string) {
	fmt.Printf("Calling command silently: %s %s\n", command, strings.Join(args, " "))
}

// confirm prompts the user for a yes/no response.
func confirm(ctx console.Context, prompt string) bool {
	response, _ := ctx.Ask(prompt + " (y/n): ")
	return strings.ToLower(response) == "y"
}

// openBrowser opens the default browser to the specified URL.
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("start", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
	}
}
