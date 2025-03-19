package tools

import "github.com/goravel/framework/contracts/console"

type PackageCommands struct {
	Commands        []console.Command
	ConsoleCommands []console.Command
	InstallCommand  func()
}

// AddCommand registers a new console command.
func (c *PackageCommands) AddCommand(command console.Command) {
	c.Commands = append(c.Commands, command)
}

// AddCommands registers multiple console commands.
func (c *PackageCommands) AddCommands(commands ...console.Command) *PackageCommands {
	c.Commands = append(c.Commands, commands...)
	return c
}

// AddConsoleCommand registers a single console command.
func (c *PackageCommands) AddConsoleCommand(command console.Command) *PackageCommands {
	c.ConsoleCommands = append(c.ConsoleCommands, command)
	return c
}

// AddConsoleCommands registers multiple console commands.
func (c *PackageCommands) AddConsoleCommands(commands ...console.Command) *PackageCommands {
	c.ConsoleCommands = append(c.ConsoleCommands, commands...)
	return c
}

// AddInstallCommand sets the installation command.
func (c *PackageCommands) AddInstallCommand(command func()) *PackageCommands {
	c.InstallCommand = command
	return c
}

// ExecuteInstallCommand runs the installation command if set.
func (c *PackageCommands) ExecuteInstallCommand() *PackageCommands {
	if c.InstallCommand != nil {
		c.InstallCommand()
	}
	return c
}
