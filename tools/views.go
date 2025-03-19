package tools

type PackageViews struct {
	Enable    bool
	Namespace *string
	ShortName string
}

// HasViews enables views and sets an optional namespace.
func (v *PackageViews) HasViews(namespace *string) *PackageViews {
	v.Enable = true
	v.Namespace = namespace
	return v
}

// ViewNamespace returns the set namespace or falls back to ShortName().
func (v *PackageViews) ViewNamespace() string {
	if v.Namespace != nil {
		return *v.Namespace
	}
	return v.ShortName
}
