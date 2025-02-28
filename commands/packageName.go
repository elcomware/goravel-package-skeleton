package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type PackageName struct{}

func NewPackageName() *PackageName {
	return &PackageName{}
}

// Signature The name and signature of the console command.
func (receiver *PackageName) Signature() string {
	return ":package_slug"
}

// Description The console command description.
func (receiver *PackageName) Description() string {
	return ":package_slug"
}

// Extend The console command extend.
func (receiver *PackageName) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *PackageName) Handle(ctx console.Context) error {
	fmt.Println("Run :package_slug command")

	return nil
}
