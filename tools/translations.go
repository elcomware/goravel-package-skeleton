package tools

type TranslationTools struct {
	Enabled bool
}

// EnableTranslation enables translations.
func (t *TranslationTools) EnableTranslation() *TranslationTools {
	t.Enabled = true
	return t
}
