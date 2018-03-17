package validator

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
)

//SpecificationParam Represents a parameter of a bioschemas spec
type SpecificationParam struct {
	Property             string `yaml:"name"`
	ExpectedType         string `yaml:"expected_type"`
	Description          string `yaml:"bsc_dec"`
	BscDescription       string `yaml:"sdo_desc"`
	Marginality          string `yaml:"marginality"`
	Cardinality          string `yaml:"cardinality"`
	ControlledVocabulary string `yaml:"controlled_vocab"`
}

//Specification A Bioschemas specification
type Specification struct {
	Name                string               `yaml:"name"`
	SpecificationParams []SpecificationParam `yaml:"properties"`
}

var (
	//Specifications All the parsed specifications
	Specifications []Specification
)

//Validate - Starts Validation
func Validate(f, u string) {
	if f != "" {
		log.Info("Validating a File " + f)
		LoadFile(f)
	} else if u != "" {
		log.Info("Validating an URL " + u)
		LoadURL(u)
	} else {
		log.Warn("Please enter valid shit")
	}
}

func vali() {
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
