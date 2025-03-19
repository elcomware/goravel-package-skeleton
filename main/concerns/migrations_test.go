package concerns

import (
	"github.com/vendorName/packageName/tools"
	"reflect"
	"testing"
)

func TestHasMigrations_HasMigration(t *testing.T) {
	type fields struct {
		runsMigrations      bool
		discoversMigrations bool
		migrationsPath      *string
		migrationFileNames  []string
	}
	type args struct {
		migrationFileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tools.HasMigrations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &tools.HasMigrations{
				runsMigrations:      tt.fields.runsMigrations,
				discoversMigrations: tt.fields.discoversMigrations,
				migrationsPath:      tt.fields.migrationsPath,
				migrationFileNames:  tt.fields.migrationFileNames,
			}
			if got := h.HasMigration(tt.args.migrationFileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasMigration() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestHasMigrations_DiscoversMigrations(t *testing.T) {
	type fields struct {
		runsMigrations      bool
		discoversMigrations bool
		migrationsPath      *string
		migrationFileNames  []string
	}
	type args struct {
		discoversMigrations bool
		path                string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tools.HasMigrations
	}{
		{
			name:   "",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &tools.HasMigrations{
				RunsMigrations:      tt.fields.runsMigrations,
				discoversMigrations: tt.fields.discoversMigrations,
				migrationsPath:      tt.fields.migrationsPath,
				migrationFileNames:  tt.fields.migrationFileNames,
			}
			if got := h.DiscoversMigrations(tt.args.discoversMigrations, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscoversMigrations() = %v, want %v", got, tt.want)
			}
		})
	}
}



func TestHasMigrations_HasMigrations(t *testing.T) {
	type fields struct {
		runsMigrations      bool
		discoversMigrations bool
		migrationsPath      *string
		migrationFileNames  []string
	}
	type args struct {
		migrationFileNames []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tools.HasMigrations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &tools.HasMigrations{
				RunsMigrations:      tt.fields.runsMigrations,
				discoversMigrations: tt.fields.discoversMigrations,
				migrationsPath:      tt.fields.migrationsPath,
				migrationFileNames:  tt.fields.migrationFileNames,
			}
			if got := h.HasMigrations(tt.args.migrationFileNames...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasMigrations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasMigrations_RunsMigrations(t *testing.T) {
	type fields struct {
		runsMigrations      bool
		discoversMigrations bool
		migrationsPath      *string
		migrationFileNames  []string
	}
	type args struct {
		runsMigrations bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *tools.HasMigrations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &tools.HasMigrations{
				RunsMigrations:      tt.fields.runsMigrations,
				discoversMigrations: tt.fields.discoversMigrations,
				migrationsPath:      tt.fields.migrationsPath,
				migrationFileNames:  tt.fields.migrationFileNames,
			}
			if got := h.RunsMigrations(tt.args.runsMigrations); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunsMigrations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasMigrations_String(t *testing.T) {
	type fields struct {
		runsMigrations      bool
		discoversMigrations bool
		migrationsPath      *string
		migrationFileNames  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &tools.HasMigrations{
				RunsMigrations:      tt.fields.runsMigrations,
				discoversMigrations: tt.fields.discoversMigrations,
				migrationsPath:      tt.fields.migrationsPath,
				migrationFileNames:  tt.fields.migrationFileNames,
			}
			if got := h.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
