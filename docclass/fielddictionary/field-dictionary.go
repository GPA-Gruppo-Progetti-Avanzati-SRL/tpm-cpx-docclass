package fielddictionary

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
)

type Mapping struct {
	Campo           string `mapstructure:"campo" yaml:"campo"`
	DocumentMapping string `mapstructure:"documentMapping" yaml:"documentMapping"`
}

type FieldDictionary map[string]Mapping

var fieldDictionary FieldDictionary

func GetDictionary() (FieldDictionary, error) {

	const semLogContext = "field-dictionary::get-dictionary"

	if fieldDictionary == nil {
		err := errors.New("field-dictionary not initialized")
		log.Error().Err(err).Msg(semLogContext)
		return nil, err
	}

	return fieldDictionary, nil
}

func GetMapping(fn string) (Mapping, error) {

	const semLogContext = "field-dictionary::get-mapping"

	var err error
	if fieldDictionary == nil {
		err = errors.New("field-dictionary not initialized")
		log.Error().Err(err).Msg(semLogContext)
		return Mapping{}, err
	}

	if fm, ok := fieldDictionary[fn]; ok {
		return fm, nil
	}

	return Mapping{}, fmt.Errorf("field not found %s", fn)
}
