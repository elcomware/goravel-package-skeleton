package tools

import (
	"github.com/goravel/framework/mocks/console"
	"github.com/stretchr/testify/assert"
	"github.com/vendorName/packageName/tools"
	"testing"
)

func TestPackageCommands_AddCommand(t *testing.T) {

	// Given
	c := tools.NewPackageCommands()
	cmd := console.NewCommand(t)
	check := assert.New(t)

	// When
	c.AddCommand(cmd)

	// Then
	check.NotNil(c.Commands)
	check.Equal(c.Commands[0], cmd)

}

func TestPackageCommands_AddCommands(t *testing.T) {
	// Given
	c := tools.NewPackageCommands()
	cmd1 := console.NewCommand(t)
	cmd2 := console.NewCommand(t)
	check := assert.New(t)

	// When
	c.AddCommands(cmd1, cmd2)

	// Then
	check.NotNil(c.Commands)
	check.Equal(c.Commands[0], cmd1)
	check.Equal(len(c.Commands), 2)
}

func TestPackageCommands_InstallCommand(t *testing.T) {
	//Given
	c := tools.NewPackageCommands()
	cmd := func() any {
		return "abc"
	}
	check := assert.New(t)

	//When
	c.AddInstallCommand(cmd)

	//Then
	check.NotNil(c.InstallCommand)
	check.Equal(c.InstallCommand(), cmd())

}

func TestPackageCommands_ExecuteInstallCommand(t *testing.T) {

	//Given
	c := tools.NewPackageCommands()
	cmd := func() any {
		return "abc"
	}
	check := assert.New(t)
	c.AddInstallCommand(cmd)

	//When
	c.ExecuteInstallCommand()

	//Then
	check.NotNil(c.InstallCommand)
	check.Equal(c.InstallCommand(), cmd())

}
