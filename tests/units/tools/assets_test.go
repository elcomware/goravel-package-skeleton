package tools

import (
	"github.com/vendorName/packageName/tools"
	"reflect"
	"testing"
)

func TestPackageAssets_EnableAssets(t *testing.T) {
	type fields struct {
		Enabled bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *tools.AssetTools
	}{
		{"can enable", fields{Enabled: true}, &tools.AssetTools{Enabled: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &tools.AssetTools{
				Enabled: tt.fields.Enabled,
			}
			if got := a.EnableAssets(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnableAssets() = %v, want %v", got, tt.want)
			}
		})
	}
}
