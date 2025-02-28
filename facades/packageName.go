package facades

import (
	"github.com/vendorName/packageName"
	"github.com/vendorName/packageName/contracts"
	"log"
)

func PackageName() contracts.PackageName {

	instance, err := packageName.App.Make(packageName.Binding)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(contracts.PackageName)
}
