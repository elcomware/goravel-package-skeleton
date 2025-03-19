package tools

type PackageViewComposers struct {
	ViewComposers map[string]string
}

/*/ NewHasViewComposers initializes the struct with an empty map.
func NewHasViewComposers() *PackageViewComposers {
	return &PackageViewComposers{ViewComposers: make(map[string]string)}
}*/

// HasViewComposer registers a view composer for one or multiple views.
func (h *PackageViewComposers) HasViewComposer(views interface{}, viewComposer string) *PackageViewComposers {
	switch v := views.(type) {
	case string:
		h.ViewComposers[v] = viewComposer
	case []string:
		for _, view := range v {
			h.ViewComposers[view] = viewComposer
		}
	}
	return h
}
