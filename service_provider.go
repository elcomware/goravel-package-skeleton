package packageName

import (
	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"
	"github.com/vendorName/packageName/commands"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

const Binding = "packageName"

//var App foundation.Application

type ServiceProvider struct {
	packageInstance *PackageName
}

// ConfigurePackage is an abstract method that must be implemented by the package.
func (sp *ServiceProvider) ConfigurePackage() {
	// This method must be implemented by the package.
}

// Register is called when the package is registered.
func (sp *ServiceProvider) Register(app foundation.Application) {
	//App = app
	sp.registeringPackage()
	//initialise package new instance and inject dependencies
	sp.packageInstance = &PackageName{} //NewPackageName(app.MakeConfig())

	//bind faced to package
	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return sp.packageInstance, nil
	})

	//Set package base directory
	sp.packageInstance.SetBasePath(sp.GetPackageBaseDir())

	sp.ConfigurePackage()

	if sp.packageInstance.FullName == "" {
		panic("Package FullName is required")
	}

	sp.registerConfigs()
	sp.packageRegistered()

}

// GetPackageBaseDir returns the base directory of the package.
func (sp *ServiceProvider) GetPackageBaseDir() string {
	reflector := reflect.ValueOf(sp).Elem()
	packagePath := reflector.Type().PkgPath()
	return filepath.Dir(packagePath)
}

// RegisteringPackage is called before the package is registered.
func (sp *ServiceProvider) registeringPackage() {
	// Optional: Override this method in your package.
}

// RegisterConfigs registers the package's configuration files.
func (sp *ServiceProvider) registerConfigs() {
	if len(sp.packageInstance.ConfigTools.ConfigFiles) == 0 {
		return
	}

	for _, configFileName := range sp.packageInstance.ConfigTools.ConfigFiles {
		configPath := filepath.Join(sp.packageInstance.GetBasePath(), "..", "config", configFileName+".go")
		facades.Config().Add(configFileName, configPath)
	}
}

// PackageRegistered is called after the package is registered.
func (sp *ServiceProvider) packageRegistered() {
	// Optional: Override this method in your package.
}

// Boot is called when the package is booted.
func (sp *ServiceProvider) Boot(app foundation.Application) {

	sp.BootingPackage()

	packageName := filepath.Join("github.com", "vendorName", "packageName")

	sp.bootPackageAssets(app, packageName).
		bootPackageCommands(app).
		bootPackageConsoleCommands(app).
		bootPackageConfigs(app, packageName).
		bootPackageMigrations(app, packageName).
		bootPackageProviders(app, packageName).
		bootPackageRoutes(app, packageName).
		bootPackageTranslations(app, packageName).
		bootPackageViews(app, packageName).
		bootPackageViewComponents(app, packageName).
		bootPackageViewComposers(app, packageName).
		bootPackageViewSharedData()

	sp.PackageBooted()

}

// BootingPackage is called before the package is booted.
func (sp *ServiceProvider) BootingPackage() {
	// Optional: Override this method in your package.
}

// PackageBooted is called after the package is booted.
func (sp *ServiceProvider) PackageBooted() {
	// Optional: Override this method in your package.
}

// BootPackageAssets boots the package's assets.
func (sp *ServiceProvider) bootPackageAssets(app foundation.Application, packageName string) *ServiceProvider {

	if !sp.packageInstance.AssetTools.Enabled || !sp.packageInstance.IsRunningInConsole() {
		return sp
	}

	//vendor and app assets paths
	vendorAssets := "public" //filepath.Join(sp.packageInstance.BasePath, "..", "resources", "dist")
	appAssets := "vendor"    //filepath.Join("vendor", sp.packageInstance.ShortName())

	// publish vendor assets to app
	app.Publishes(packageName, map[string]string{
		vendorAssets: app.PublicPath(appAssets),
	}, "packageName-assets")

	return sp
}

// BootPackageCommands boots the package's commands.
func (sp *ServiceProvider) bootPackageCommands(app foundation.Application) *ServiceProvider {

	// register package command
	app.Commands(
		[]console.Command{
			commands.NewPackageName(),
			commands.NewInstallCommand(commands.SudoPackage{
				ShortName:               sp.packageInstance.ShortName,
				FullName:                sp.packageInstance.FullName,
				PublishableProviderName: sp.packageInstance.PublishableProviderName,
			}),
		})

	if len(sp.packageInstance.CommandTools.Commands) == 0 {
		return sp
	}
	///register other commands if available
	//facades.Artisan().Register(sp.packageInstance.commands)
	app.Commands(
		sp.packageInstance.CommandTools.Commands,
	)

	return sp
}

// BootPackageConsoleCommands boots the package's console commands.
func (sp *ServiceProvider) bootPackageConsoleCommands(app foundation.Application) *ServiceProvider {
	if len(sp.packageInstance.CommandTools.ConsoleCommands) == 0 {
		return sp
	}

	app.Commands(
		sp.packageInstance.CommandTools.ConsoleCommands,
	)

	return sp
}

// BootPackageConfigs boots the package's configuration files.
func (sp *ServiceProvider) bootPackageConfigs(app foundation.Application, packageName string) *ServiceProvider {
	if !sp.packageInstance.IsRunningInConsole() {
		return sp
	}

	//publish package configs
	app.Publishes(packageName, map[string]string{
		"config/packageName.go": app.ConfigPath("packageName.go"),
	}, "packageName-config")

	// publish other configs
	for _, configFileName := range sp.packageInstance.ConfigTools.ConfigFiles {
		//vendorConfig := filepath.Join(sp.packageInstance.BasePath(), "..", "config", configFileName+".go")
		vendorConfig := filepath.Join("config", configFileName+".go")
		appConfig := filepath.Join(configFileName + ".go")

		app.Publishes(packageName, map[string]string{
			vendorConfig: app.ConfigPath(appConfig),
		}, "all-configs")
	}

	return sp
}

// BootPackageMigrations boots the package's migrations.
func (sp *ServiceProvider) bootPackageMigrations(app foundation.Application, packageName string) *ServiceProvider {
	if sp.packageInstance.MigrationTools.Discovers {
		sp.discoverPackageMigrations(app, packageName)
		return sp
	}

	now := time.Now()

	for _, migrationFileName := range sp.packageInstance.MigrationTools.Files {
		packageName := filepath.Join("github.com", "vendorName", "packageName")
		//vendorMigration := filepath.Join(sp.packageInstance.BasePath, "..", "database", "migrations", migrationFileName+".go")
		vendorMigration := filepath.Join("database", "migrations", migrationFileName+".go")
		appMigration := sp.GenerateMigrationName(migrationFileName, now.Add(time.Second))

		if !fileExists(vendorMigration) {
			vendorMigration += ".stub"
		}

		if sp.packageInstance.IsRunningInConsole() {
			app.Publishes(packageName, map[string]string{
				vendorMigration: app.DatabasePath(appMigration),
			}, "migrations")

		}

		/*if sp.packageInstance.CanRun {
			app.DatabasePath(appMigration)
			facades.Migration.Load(vendorMigration)
			app.DatabasePath()
		}*/
	}

	return sp
}

// BootPackageProviders boots the package's service providers.
func (sp *ServiceProvider) bootPackageProviders(app foundation.Application, packageName string) *ServiceProvider {
	if sp.packageInstance.PublishableProviderName == "" || !sp.packageInstance.IsRunningInConsole() {
		return sp
	}

	providerName := sp.packageInstance.PublishableProviderName
	vendorProvider := filepath.Join(sp.packageInstance.BasePath, "stubs", "providers", providerName+".go.stub")
	appProvider := filepath.Join(app.Path(), "Providers", providerName+"_service_provider.go")

	app.Publishes(packageName, map[string]string{
		vendorProvider: appProvider,
	}, "providers")

	return sp

}

// BootPackageRoutes boots the package's routes.
func (sp *ServiceProvider) bootPackageRoutes(app foundation.Application, packageName string) *ServiceProvider {
	if len(sp.packageInstance.RouteTools.FileNames) == 0 {
		return sp
	}

	for _, routeFileName := range sp.packageInstance.RouteTools.FileNames {
		vendorRoutes := filepath.Join(sp.packageInstance.BasePath, "..", "routes", routeFileName+".php")
		appRoute := filepath.Join(app.Path(), "routes", packageName+"_service_provider.go")

		app.Publishes(packageName, map[string]string{
			vendorRoutes: appRoute,
		}, "routes")
	}

	return sp
}

// BootPackageTranslations boots the package's translations.
func (sp *ServiceProvider) bootPackageTranslations(app foundation.Application, packageName string) *ServiceProvider {
	if !sp.packageInstance.TranslationTools.Enabled {
		return sp
	}

	vendorTranslations := filepath.Join(sp.packageInstance.BasePath, "..", "resources", "lang")
	appTranslations := filepath.Join("resources", "lang", "vendor", sp.packageInstance.ShortName)

	/*facades.Translation.Load(vendorTranslations, sp.packageInstance.ShortName)
	facades.Translation.LoadJSON(vendorTranslations)
	facades.Translation.LoadJSON(appTranslations)*/

	if sp.packageInstance.IsRunningInConsole() {

		app.Publishes(packageName, map[string]string{
			vendorTranslations: appTranslations,
		}, "translations")
	}

	return sp
}

// BootPackageViews boots the package's views.
func (sp *ServiceProvider) bootPackageViews(app foundation.Application, packageName string) *ServiceProvider {
	if !sp.packageInstance.ViewsTools.Enabled {
		return sp
	}

	//namespace := sp.packageInstance.ViewsTools.ViewNamespace()
	vendorViews := filepath.Join(sp.packageInstance.BasePath, "..", "resources", "views")
	appViews := filepath.Join("resources", "views", "vendor", sp.packageInstance.ShortName)

	//facades.View.Load(vendorViews, namespace)

	if sp.packageInstance.IsRunningInConsole() {
		app.Publishes(packageName, map[string]string{
			vendorViews: appViews,
		})
	}

	return sp
}

// BootPackageViewComponents boots the package's view components.
func (sp *ServiceProvider) bootPackageViewComponents(app foundation.Application, packageName string) *ServiceProvider {
	if len(sp.packageInstance.ViewComponentTools.Components) == 0 {
		return sp
	}

	/*for componentClass, prefix := range sp.packageInstance.ViewComponentTools.Components {
		facades.View.Component(prefix, componentClass)
	}*/

	if sp.packageInstance.IsRunningInConsole() {
		vendorComponents := filepath.Join(sp.packageInstance.BasePath, "Components")
		appComponents := filepath.Join("app", "View", "Components", "vendor", sp.packageInstance.ShortName)

		//facades.Publisher.Publish(vendorComponents, appComponents, sp.packageInstance.FullName+"-components")
		app.Publishes(packageName, map[string]string{
			vendorComponents: appComponents,
		})
	}

	return sp
}

// BootPackageViewComposers boots the package's view composers.
func (sp *ServiceProvider) bootPackageViewComposers(app foundation.Application, packageName string) *ServiceProvider {
	if len(sp.packageInstance.ViewComposerTools.Composers) == 0 {
		return sp
	}

	for viewName, viewComposer := range sp.packageInstance.ViewComposerTools.Composers {
		app.Publishes(packageName, map[string]string{
			viewName: viewComposer,
		})
	}

	return sp
}

// BootPackageViewSharedData boots the package's shared view data.
func (sp *ServiceProvider) bootPackageViewSharedData() *ServiceProvider {
	if len(sp.packageInstance.ViewSharedDataTools.SharedData) == 0 {
		return sp
	}

	/*for FullName, value := range sp.packageInstance.ViewSharedDataTools.SharedData {
		facades.View.Share(FullName, value)
	}*/

	return sp
}

// DiscoverPackageMigrations discovers and registers package migrations.
func (sp *ServiceProvider) discoverPackageMigrations(app foundation.Application, packageName string) {
	now := time.Now()

	migrationsPath := strings.Trim(*sp.packageInstance.MigrationTools.Path, "/")

	files, _ := os.ReadDir(filepath.Join(sp.packageInstance.BasePath, "..", migrationsPath))

	for _, file := range files {
		vendorMigrations := filepath.Join(sp.packageInstance.BasePath, "..", migrationsPath, file.Name())
		migrationFileName := strings.TrimSuffix(file.Name(), ".go")
		migrationFileName = strings.TrimSuffix(migrationFileName, ".stub")

		appMigration := sp.GenerateMigrationName(migrationFileName, now.Add(time.Second))

		if sp.packageInstance.IsRunningInConsole() {
			//facades.Publisher.Publish(filePath, appMigration, sp.packageInstance.ShortName()+"-migrations")
			app.Publishes(packageName, map[string]string{
				vendorMigrations: appMigration,
			})
		}

	}
}

// GenerateMigrationName generates a unique migration FullName.
func (sp *ServiceProvider) GenerateMigrationName(migrationFileName string, now time.Time) string {
	migrationsPath := filepath.Join("database", "migrations", filepath.Dir(migrationFileName))
	migrationFileName = filepath.Base(migrationFileName)

	//len := len(migrationFileName) + 4

	if strings.Contains(migrationFileName, "/") {
		migrationsPath = filepath.Join(migrationsPath, strings.Split(migrationFileName, "/")[0])
		migrationFileName = strings.Split(migrationFileName, "/")[1]
	}

	matches, _ := filepath.Glob(filepath.Join(migrationsPath, "*.go"))
	for _, match := range matches {
		if strings.HasSuffix(match, migrationFileName+".go") {
			return match
		}
	}

	timestamp := now.Format("2006_01_02_150405")
	migrationFileName = strings.ReplaceAll(migrationFileName, "-", "_") + ".go"

	return filepath.Join(migrationsPath, timestamp+"_"+migrationFileName)
}

// FileExists checks if a file exists.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
