package units

import (
	"github.com/stretchr/testify/assert"
	"github.com/vendorName/packageName"
	"testing"
	"time"
)

func TestServiceProvider_GenerateMigrationName(t *testing.T) {

	//Given
	sp := packageName.ServiceProvider{}
	migrationFile := "user"
	now := time.Now()
	ck := assert.New(t)

	//When
	actual := sp.GenerateMigrationName(migrationFile, now)
	//fmt.Println(name)

	//Then
	expect := "database\\migrations\\" + now.Format("2006_01_02_150405") + "_" + migrationFile + ".go"
	ck.Equal(expect, actual)

}

func TestServiceProvider_GetPackageBaseDir(t *testing.T) {
	//Given
	sp := packageName.ServiceProvider{}

	//When
	actual := sp.GetPackageBaseDir()

	//Then
	expect := "github.com\\vendorName"
	ck := assert.New(t)
	ck.NotNil(actual)
	ck.Equal(expect, actual)

}
