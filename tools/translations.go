package tools

type PackageTranslations struct {
	Enable bool
}

// EnableTranslation enables translations.
func (t *PackageTranslations) EnableTranslation() *PackageTranslations {
	t.Enable = true
	return t
}
