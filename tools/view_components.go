package tools

type PackageViewComponents struct {
	ViewComponents map[string]string

	HasViews      bool
	ViewNamespace *string
}

/*/ NewViewsComponentsService initializes the struct with an empty map.
func NewViewsComponentsService() *PackageViewComponents {
	return &PackageViewComponents{ViewComponents: make(map[string]string)}
}*/

// AddViewComponent adds a single view component with a prefix.
func (c *PackageViewComponents) AddViewComponent(prefix, viewComponentName string) *PackageViewComponents {
	c.ViewComponents[viewComponentName] = prefix
	return c
}

// AddViewComponents adds multiple view components with the same prefix.
func (c *PackageViewComponents) AddViewComponents(prefix string, viewComponentNames ...string) *PackageViewComponents {
	for _, componentName := range viewComponentNames {
		c.ViewComponents[componentName] = prefix
	}
	return c
}
