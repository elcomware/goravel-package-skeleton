package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type PackageNameCommand struct{}

func NewPackageName() *PackageNameCommand {
	return &PackageNameCommand{}
}

// Signature The name and signature of the console command.
func (receiver *PackageNameCommand) Signature() string {
	return ":packageName"
}

// Description The console command description.
func (receiver *PackageNameCommand) Description() string {
	return ":packageName command"
}

// Extend The console command extend.
func (receiver *PackageNameCommand) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *PackageNameCommand) Handle(ctx console.Context) error {

	fmt.Println("Run :packageName command")

	return nil
}
