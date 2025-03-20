package tools

type AssetTools struct {
	Enabled bool
}

func (a *AssetTools) EnableAssets() *AssetTools {
	a.Enabled = true
	return a
}
