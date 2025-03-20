package tools

import "github.com/goravel/framework/facades"

type ViewSharedDataTools struct {
	SharedData map[string]interface{}
}

/*/ NewHasViewSharedData initializes the struct with an empty map.
func NewHasViewSharedData() *ViewSharedDataTools {
	return &ViewSharedDataTools{SharedData: make(map[string]interface{})}
}*/

// SharesDataWithAllViews shares data with all views.
func (h *ViewSharedDataTools) SharesDataWithAllViews(name string, value interface{}) *ViewSharedDataTools {
	h.SharedData[name] = value
	return h
}

// ShareAllDataWithViews applies shared data to all views in Goravel.
func (h *ViewSharedDataTools) ShareAllDataWithViews() {
	for name, value := range h.SharedData {
		facades.View().Share(name, value)
	}
}
