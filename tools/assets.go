package tools

type PackageAssets struct {
	Enabled bool
}

func (a *PackageAssets) EnableAssets() *PackageAssets {
	a.Enabled = true
	return a
}
