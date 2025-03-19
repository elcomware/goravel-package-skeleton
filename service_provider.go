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

var App foundation.Application

type ServiceProvider struct {
	packageInstance *PackageName
}

// ConfigurePackage is an abstract method that must be implemented by the package.
func (sp *ServiceProvider) ConfigurePackage(pkg *PackageName) {
	// This method must be implemented by the package.
}

// Register is called when the package is registered.
func (sp *ServiceProvider) Register(app foundation.Application) {
	App = app
	sp.registeringPackage(app)
	//initialise package new instance and inject dependencies
	sp.packageInstance = NewPackageName(app.MakeConfig())

	//bind faced to package
	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return sp.packageInstance, nil
	})

	//Set package base directory
	sp.packageInstance.SetBasePath(sp.getPackageBaseDir())

	sp.ConfigurePackage(sp.packageInstance)

	if sp.packageInstance.name == "" {
		panic("Package name is required")
	}

	sp.registerConfigs()
	sp.packageRegistered()

}

// GetPackageBaseDir returns the base directory of the package.
func (sp *ServiceProvider) getPackageBaseDir() string {
	reflector := reflect.ValueOf(sp).Elem()
	packagePath := reflector.Type().PkgPath()
	return filepath.Dir(packagePath)
}

// RegisteringPackage is called before the package is registered.
func (sp *ServiceProvider) registeringPackage(app foundation.Application) {
	// Optional: Override this method in your package.
}

// RegisterConfigs registers the package's configuration files.
func (sp *ServiceProvider) registerConfigs() {
	if len(sp.packageInstance.configFileNames) == 0 {
		return
	}

	for _, configFileName := range sp.packageInstance.configFileNames {
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
	app.Publishes("github.com/vendorName/packageName", map[string]string{
		"config/packageName.go": app.ConfigPath("packageName.go"),
	}, "packageName-config")
	app.Commands(
		[]console.Command{
			commands.NewPackageName(),
		})

	sp.bootPackageAssets(app).
		bootPackageCommands(app).
		bootPackageConsoleCommands(app).
		bootPackageConfigs(app).
		bootPackageMigrations(app).
		bootPackageProviders(app).
		bootPackageRoutes(app).
		bootPackageTranslations(app).
		bootPackageViews(app).
		bootPackageViewComponents(app).
		bootPackageViewComposers(app).
		bootPackageViewSharedData(app)

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
func (sp *ServiceProvider) bootPackageAssets(app foundation.Application) *ServiceProvider {

	if !sp.packageInstance.hasAssets || !sp.packageInstance.IsRunningInConsole() {
		return sp
	}

	//vendor and app assets paths
	vendorAssets := "public" //filepath.Join(sp.packageInstance.basePath, "..", "resources", "dist")
	appAssets := "vendor"    //filepath.Join("vendor", sp.packageInstance.ShortName())

	// publish vendor assets to app
	app.Publishes("github.com/goravel/example-package", map[string]string{
		vendorAssets: app.PublicPath(appAssets),
	}, "packageName-assets")

	return sp
}

// BootPackageCommands boots the package's commands.
func (sp *ServiceProvider) bootPackageCommands(app foundation.Application) *ServiceProvider {
	if len(sp.packageInstance.commands) == 0 {
		return sp
	}

	//facades.Artisan().Register(sp.packageInstance.commands)
	app.Commands(
		sp.packageInstance.commands,
	)

	return sp
}

// BootPackageConsoleCommands boots the package's console commands.
func (sp *ServiceProvider) bootPackageConsoleCommands(app foundation.Application) *ServiceProvider {
	if len(sp.packageInstance.ConsoleCommands) == 0 {
		return sp
	}

	app.Commands(
		sp.packageInstance.ConsoleCommands,
	)

	return sp
}

// BootPackageConfigs boots the package's configuration files.
func (sp *ServiceProvider) bootPackageConfigs(app foundation.Application) *ServiceProvider {
	if !sp.packageInstance.IsRunningInConsole() {
		return sp
	}

	for _, configFileName := range sp.packageInstance.configFileNames {
		//vendorConfig := filepath.Join(sp.packageInstance.basePath(), "..", "config", configFileName+".go")
		packageName := filepath.Join("github.com", "vendorName", "packageName")
		vendorConfig := filepath.Join("config", configFileName+".go")
		appConfig := filepath.Join(configFileName + ".go")

		app.Publishes(packageName, map[string]string{
			vendorConfig: app.ConfigPath(appConfig),
		}, "all-configs")
	}

	return sp
}

// BootPackageMigrations boots the package's migrations.
func (sp *ServiceProvider) bootPackageMigrations(app foundation.Application) *ServiceProvider {
	if sp.packageInstance.discoverMigrations {
		sp.discoverPackageMigrations()
		return sp
	}

	now := time.Now()

	for _, migrationFileName := range sp.packageInstance.migrationFileNames {
		packageName := filepath.Join("github.com", "vendorName", "packageName")
		//vendorMigration := filepath.Join(sp.packageInstance.basePath, "..", "database", "migrations", migrationFileName+".go")
		vendorMigration := filepath.Join("database", "migrations", migrationFileName+".go")
		appMigration := sp.generateMigrationName(migrationFileName, now.Add(time.Second))

		if !fileExists(vendorMigration) {
			vendorMigration += ".stub"
		}

		if sp.packageInstance.IsRunningInConsole() {
			app.Publishes(packageName, map[string]string{
				vendorMigration: app.DatabasePath(appMigration),
			}, "migrations")

		}

		if sp.packageInstance.RunsMigrations {
			facades.Migration.Load(vendorMigration)
			app.DatabasePath()
		}
	}

	return sp
}

// BootPackageProviders boots the package's service providers.
func (sp *ServiceProvider) bootPackageProviders() *ServiceProvider {
	if sp.packageInstance.PublishableProviderName == "" || !facades.App.RunningInConsole() {
		return sp
	}

	providerName := sp.packageInstance.PublishableProviderName
	vendorProvider := filepath.Join(sp.packageInstance.basePath(), "..", "resources", "stubs", providerName+".php.stub")
	appProvider := filepath.Join("app", "Providers", providerName+".php")

	facades.Publisher.Publish(vendorProvider, appProvider, sp.packageInstance.ShortName()+"-provider")

	return sp
}

// BootPackageRoutes boots the package's routes.
func (sp *ServiceProvider) bootPackageRoutes() *ServiceProvider {
	if len(sp.packageInstance.RouteFileNames) == 0 {
		return sp
	}

	for _, routeFileName := range sp.packageInstance.RouteFileNames {
		routePath := filepath.Join(sp.packageInstance.basePath(), "..", "routes", routeFileName+".php")
		facades.Route.Load(routePath)
	}

	return sp
}

// BootPackageTranslations boots the package's translations.
func (sp *ServiceProvider) bootPackageTranslations() *ServiceProvider {
	if !sp.packageInstance.HasTranslations {
		return sp
	}

	vendorTranslations := filepath.Join(sp.packageInstance.basePath(), "..", "resources", "lang")
	appTranslations := filepath.Join("resources", "lang", "vendor", sp.packageInstance.ShortName())

	facades.Translation.Load(vendorTranslations, sp.packageInstance.ShortName())
	facades.Translation.LoadJSON(vendorTranslations)
	facades.Translation.LoadJSON(appTranslations)

	if facades.App.RunningInConsole() {
		facades.Publisher.Publish(vendorTranslations, appTranslations, sp.packageInstance.ShortName()+"-translations")
	}

	return sp
}

// BootPackageViews boots the package's views.
func (sp *ServiceProvider) bootPackageViews(app foundation.Application) *ServiceProvider {
	if !sp.packageInstance.HasViews {
		return sp
	}

	namespace := sp.packageInstance.ViewNamespace
	vendorViews := filepath.Join(sp.packageInstance.basePath(), "..", "resources", "views")
	appViews := filepath.Join("resources", "views", "vendor", sp.packageInstance.ShortName())

	facades.View.Load(vendorViews, namespace)

	if sp.packageInstance.IsRunningInConsole() {
		facades.Publisher.Publish(vendorViews, appViews, sp.packageInstance.ShortName()+"-views")
		app.Publishes()
	}

	return sp
}

// BootPackageViewComponents boots the package's view components.
func (sp *ServiceProvider) bootPackageViewComponents() *ServiceProvider {
	if len(sp.packageInstance.ViewComponents) == 0 {
		return sp
	}

	for componentClass, prefix := range sp.packageInstance.ViewComponents {
		facades.View.Component(prefix, componentClass)
	}

	if facades.App.RunningInConsole() {
		vendorComponents := filepath.Join(sp.packageInstance.basePath(), "Components")
		appComponents := filepath.Join("app", "View", "Components", "vendor", sp.packageInstance.ShortName())

		facades.Publisher.Publish(vendorComponents, appComponents, sp.packageInstance.Name+"-components")
	}

	return sp
}

// BootPackageViewComposers boots the package's view composers.
func (sp *ServiceProvider) bootPackageViewComposers() *ServiceProvider {
	if len(sp.packageInstance.ViewComposers) == 0 {
		return sp
	}

	for viewName, viewComposer := range sp.packageInstance.ViewComposers {
		facades.View.Composer(viewName, viewComposer)
	}

	return sp
}

// BootPackageViewSharedData boots the package's shared view data.
func (sp *ServiceProvider) bootPackageViewSharedData() *ServiceProvider {
	if len(sp.packageInstance.SharedViewData) == 0 {
		return sp
	}

	for name, value := range sp.packageInstance.SharedViewData {
		facades.View.Share(name, value)
	}

	return sp
}

// DiscoverPackageMigrations discovers and registers package migrations.
func (sp *ServiceProvider) discoverPackageMigrations() {
	now := time.Now()
	migrationsPath := strings.Trim(sp.packageInstance.MigrationsPath, "/")

	files, _ := os.ReadDir(filepath.Join(sp.packageInstance.basePath, "..", migrationsPath))

	for _, file := range files {
		filePath := filepath.Join(sp.packageInstance.basePath(), "..", migrationsPath, file.Name())
		migrationFileName := strings.TrimSuffix(file.Name(), ".php")
		migrationFileName = strings.TrimSuffix(migrationFileName, ".stub")

		appMigration := sp.generateMigrationName(migrationFileName, now.Add(time.Second))

		if facades.App.RunningInConsole() {
			facades.Publisher.Publish(filePath, appMigration, sp.packageInstance.ShortName()+"-migrations")
		}

		if sp.packageInstance.RunsMigrations {
			facades.Migration.Load(filePath)
		}
	}
}

// GenerateMigrationName generates a unique migration name.
func (sp *ServiceProvider) generateMigrationName(migrationFileName string, now time.Time) string {
	migrationsPath := filepath.Join("database", "migrations", filepath.Dir(migrationFileName))
	migrationFileName = filepath.Base(migrationFileName)

	len := len(migrationFileName) + 4

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
