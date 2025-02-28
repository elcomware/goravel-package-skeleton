package packageName

import (
	"github.com/elcomware/fortify/commands"
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "packageName"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewPackageName(app.MakeConfig()), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/vendorName/packageName", map[string]string{
		"config/packageName.go": app.ConfigPath("packageName.go"),
	})
	app.Commands([]console.Command{
		commands.NewpackageName(),
	})
}
