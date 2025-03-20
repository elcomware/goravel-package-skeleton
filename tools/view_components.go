package tools

type ViewComponentTools struct {
	Components map[string]string

	HasViews  bool
	Namespace *string
}

// NewPackageViewComponents initializes the struct with an empty map.
func NewPackageViewComponents() *ViewComponentTools {
	return &ViewComponentTools{Components: make(map[string]string)}
}

// AddViewComponent adds a single view component with a prefix.
func (c *ViewComponentTools) AddViewComponent(prefix, viewComponentName string) *ViewComponentTools {
	c.Components[viewComponentName] = prefix
	return c
}

// AddViewComponents adds multiple view components with the same prefix.
func (c *ViewComponentTools) AddViewComponents(prefix string, viewComponentNames ...string) *ViewComponentTools {
	for _, componentName := range viewComponentNames {
		c.Components[componentName] = prefix
	}
	return c
}
