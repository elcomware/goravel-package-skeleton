package tools

import "github.com/goravel/framework/contracts/console"

type CommandTools struct {
	Commands        []console.Command
	ConsoleCommands []console.Command
	InstallCommand  func() any
}

func NewPackageCommands() *CommandTools {
	return &CommandTools{}
}

// AddCommand registers a new console command.
func (c *CommandTools) AddCommand(command console.Command) *CommandTools {
	c.Commands = append(c.Commands, command)
	return c
}

// AddCommands registers multiple console commands.
func (c *CommandTools) AddCommands(commands ...console.Command) *CommandTools {
	c.Commands = append(c.Commands, commands...)
	return c
}

// AddConsoleCommand registers a single console command.
func (c *CommandTools) AddConsoleCommand(command console.Command) *CommandTools {
	c.ConsoleCommands = append(c.ConsoleCommands, command)
	return c
}

// AddConsoleCommands registers multiple console commands.
func (c *CommandTools) AddConsoleCommands(commands ...console.Command) *CommandTools {
	c.ConsoleCommands = append(c.ConsoleCommands, commands...)
	return c
}

// AddInstallCommand sets the installation command.
func (c *CommandTools) AddInstallCommand(command func() any) *CommandTools {
	c.InstallCommand = command
	return c
}

// ExecuteInstallCommand runs the installation command if set.
func (c *CommandTools) ExecuteInstallCommand() *CommandTools {
	if c.InstallCommand != nil {
		c.InstallCommand()
	}
	return c
}
