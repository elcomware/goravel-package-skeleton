package tools

type PackageConfigs struct {
	FileNames []string
	ShortName string
}

// AddConfigFile sets configuration file names.
func (c *PackageConfigs) AddConfigFile(configFileName ...string) *PackageConfigs {
	if len(configFileName) == 0 {
		// Default to the package's short name
		configFileName = []string{c.ShortName}
	}

	c.FileNames = append(c.FileNames, configFileName...)
	return c
}

/*func (c *PackageConfigs) HasConfigFile(configFileName ...string) *PackageConfigs {
	if len(configFileName) == 0 {
		configFileName = []string{c.ShortName}
	}

	c.FileNames = configFileName
	return c
}
*/
