package tools

import "github.com/goravel/framework/facades"

type PackageViewSharedData struct {
	SharedViewData map[string]interface{}
}

/*/ NewHasViewSharedData initializes the struct with an empty map.
func NewHasViewSharedData() *PackageViewSharedData {
	return &PackageViewSharedData{SharedViewData: make(map[string]interface{})}
}*/

// SharesDataWithAllViews shares data with all views.
func (h *PackageViewSharedData) SharesDataWithAllViews(name string, value interface{}) *PackageViewSharedData {
	h.SharedViewData[name] = value
	return h
}

// ShareAllDataWithViews applies shared data to all views in Goravel.
func (h *PackageViewSharedData) ShareAllDataWithViews() {
	for name, value := range h.SharedViewData {
		facades.View().Share(name, value)
	}
}
