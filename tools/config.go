package tools

type ConfigTools struct {
	ConfigFiles []string
	ShortName   string
}

// AddConfigFile sets configuration file names.
func (c *ConfigTools) AddConfigFile(configFileName ...string) *ConfigTools {
	if len(configFileName) == 0 {
		// Default to the package's short name
		configFileName = []string{c.ShortName}
	}

	c.ConfigFiles = append(c.ConfigFiles, configFileName...)
	return c
}

/*func (c *ConfigTools) HasConfigFile(configFileName ...string) *ConfigTools {
	if len(configFileName) == 0 {
		configFileName = []string{c.ShortName}
	}

	c.Files = configFileName
	return c
}
*/
