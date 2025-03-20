package tools

import (
	"github.com/stretchr/testify/assert"
	"github.com/vendorName/packageName/tools"
	"testing"
)

func TestViewComponent_AddViewComponent(t *testing.T) {

	//Given
	v := tools.NewPackageViewComponents()
	ck := assert.New(t)

	//When
	v.AddViewComponent("ui", "button")
	//Then
	ck.NotNil(v.Components["button"])
	ck.Equal(v.Components["button"], "ui")
}

func TestViewComponent_AddViewComponents(t *testing.T) {

	//Given
	v := tools.NewPackageViewComponents()
	ck := assert.New(t)

	//When
	v.AddViewComponents("ui", "card", "modal")

	//Then
	ck.NotNil(v.Components["card"])
}
