package tools

type ViewTools struct {
	Enabled   bool
	Namespace *string
	ShortName string
}

// HasViews enables views and sets an optional namespace.
func (v *ViewTools) HasViews(namespace *string) *ViewTools {
	v.Enabled = true
	v.Namespace = namespace
	return v
}

// ViewNamespace returns the set namespace or falls back to ShortName().
func (v *ViewTools) ViewNamespace() string {
	if v.Namespace != nil {
		return *v.Namespace
	}
	return v.ShortName
}
