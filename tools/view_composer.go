package tools

type ViewComposerTools struct {
	Composers map[string]string
}

/*/ NewHasViewComposers initializes the struct with an empty map.
func NewHasViewComposers() *ViewComposerTools {
	return &ViewComposerTools{Composers: make(map[string]string)}
}*/

// HasViewComposer registers a view composer for one or multiple views.
func (h *ViewComposerTools) HasViewComposer(views interface{}, viewComposer string) *ViewComposerTools {
	switch v := views.(type) {
	case string:
		h.Composers[v] = viewComposer
	case []string:
		for _, view := range v {
			h.Composers[view] = viewComposer
		}
	}
	return h
}
