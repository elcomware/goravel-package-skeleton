package tools

import (
	"strings"
)

// MigrationTools is a struct that encapsulates migration-related configurations.
// It provides methods to configure whether migrations should run, discover migrations,
// and manage migration file names.
type MigrationTools struct {
	// CanRun indicates whether migrations should be executed.
	CanRun bool

	// Discovers indicates whether migrations should be discovered automatically.
	Discovers bool

	// migrationsPath specifies the directory path where migrations are located.
	// It is a pointer to allow for a nullable value (similar to PHP's `null`).
	Path *string

	// migrationFileNames is a slice of migration file names to be applied.
	Files []string
}

func NewPackageMigration() *MigrationTools {
	return &MigrationTools{}
}

// DiscoversMigrations Allows discovery of migration files
func (m *MigrationTools) DiscoversMigrations() *MigrationTools {
	m.Discovers = true
	return m
}

// AddMigration appends a single migration file name to the `migrationFileNames` slice.
// It returns the receiver to enable method chaining.
func (m *MigrationTools) AddMigration(migrationFileName string) *MigrationTools {
	m.Files = append(m.Files, migrationFileName)
	return m
}

// AddMigrations  registers multiple migration files.
func (m *MigrationTools) AddMigrations(migrationFiles ...string) *MigrationTools {
	m.Files = append(m.Files, migrationFiles...)
	return m
}

// RunsMigrations  sets the `runsMigrations` flag to indicate whether migrations should run.
// It returns the receiver to enable method chaining.
func (m *MigrationTools) RunsMigrations(runsMigrations bool) *MigrationTools {
	m.CanRun = runsMigrations
	return m
}

// String provides a human-readable representation of the HasMigrations struct.
// This is useful for debugging and logging.
func (m *MigrationTools) String() string {
	var builder strings.Builder

	builder.WriteString("MigrationTools {\n")
	builder.WriteString("  RunsMigrations: ")
	builder.WriteString(func() string {
		if m.CanRun {
			return "true"
		}
		return "false"
	}())
	builder.WriteString(",\n")

	builder.WriteString("  DiscoversMigrations: ")
	builder.WriteString(func() string {
		if m.Discovers {
			return "true"
		}
		return "false"
	}())
	builder.WriteString(",\n")

	builder.WriteString("  MigrationsPath: ")
	if m.Path != nil {
		builder.WriteString(*m.Path)
	} else {
		builder.WriteString("nil")
	}
	builder.WriteString(",\n")

	builder.WriteString("  migrationFileNames: ")
	builder.WriteString(strings.Join(m.Files, ", "))
	builder.WriteString("\n}")

	return builder.String()
}
