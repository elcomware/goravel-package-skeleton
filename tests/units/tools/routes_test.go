package tools

import (
	"github.com/stretchr/testify/assert"
	"github.com/vendorName/packageName/tools"
	"testing"
)

func TestPackageRoute_AddRoute(t *testing.T) {

	// Given: I have a package route instance and file name
	r := tools.NewPackageRoutes()
	fileName := "testdata/routes.go"

	// When: I call the AddRoute method
	r.AddRoute(fileName)

	// Then: The routeFileName should not be empty
	assert.NotNil(t, r.FileNames)
	// Then: The routeFileName should not be same as
	assert.Equal(t, r.FileNames[0], fileName)

}

func TestPackageRoute_AddRoutes(t *testing.T) {

	// Given: I have a package route instance and file name
	r := tools.NewPackageRoutes()
	fileNames := []string{"path1", "path2"}

	// When: I call the AddRoute method
	r.AddRoutes(fileNames)

	// Then: The routeFileName should not be empty
	assert.NotNil(t, r.FileNames)
	// Then: The routeFileName should not be same as
	assert.Equal(t, r.FileNames[1], "path2")

}
