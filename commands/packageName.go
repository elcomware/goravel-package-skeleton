package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type packageName struct{}

func NewpackageName() *packageName {
	return &packageName{}
}

// Signature The name and signature of the console command.
func (receiver *packageName) Signature() string {
	return ":package_slug"
}

// Description The console command description.
func (receiver *packageName) Description() string {
	return ":package_slug"
}

// Extend The console command extend.
func (receiver *packageName) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *packageName) Handle(ctx console.Context) error {
	fmt.Println("Run :package_slug command")

	return nil
}
