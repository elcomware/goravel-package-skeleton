package tools

import (
	"github.com/vendorName/packageName/tools"
	"reflect"
	"testing"
)

func TestPackageConfigs_AddConfigFile(t *testing.T) {
	type fields struct {
		FileNames []string
	}
	type args struct {
		configFileName []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tools.ConfigTools
	}{
		{
			name: "adds config",
			fields: fields{
				FileNames: []string{},
			},
			args: args{configFileName: []string{"sample-goravel"}},
			want: &tools.ConfigTools{
				ConfigFiles: []string{"sample-goravel"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &tools.ConfigTools{
				ConfigFiles: tt.fields.FileNames,
			}
			if got := c.AddConfigFile(tt.args.configFileName...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
