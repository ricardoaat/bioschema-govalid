package validator

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

func Validate() {
	schemaLoader := gojsonschema.NewReferenceLoader("file:///Users/arcila/ebi/dev/bioschemas-govalid/build/event.json")
	documentLoader := gojsonschema.NewReferenceLoader("file:///Users/arcila/ebi/dev/bioschemas-govalid/build/tovalid.json")
	log.Info(schemaLoader)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		fmt.Println(err)
	}

	if result.Valid() {
		fmt.Printf("The document is valid\n")
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
	}
}
