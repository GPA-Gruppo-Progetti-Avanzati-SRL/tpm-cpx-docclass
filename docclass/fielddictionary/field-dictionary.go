package fielddictionary

import (
	_ "embed"
	"fmt"
)

//go:embed field-dictionary.yml
var fieldRegistryFile []byte

type Mapping struct {
	Campo           string `mapstructure:"campo" yaml:"campo"`
	DocumentMapping string `mapstructure:"documentMapping" yaml:"documentMapping"`
}

type FieldRegistry map[string]Mapping

var fieldRegistry FieldRegistry

func GetMapping(fn string) (Mapping, error) {

	var err error
	if fieldRegistry == nil {
		fieldRegistry, err = readDictionaryFromEmbeddedData()
		if err != nil {
			return Mapping{}, nil
		}
	}

	if fm, ok := fieldRegistry[fn]; ok {
		return fm, nil
	}

	return Mapping{}, fmt.Errorf("field not found %s", fn)
}
