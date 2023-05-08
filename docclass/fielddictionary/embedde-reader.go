package fielddictionary

import (
	"gopkg.in/yaml.v3"
	"strings"
)

func readDictionaryFromEmbeddedData() (FieldRegistry, error) {

	fm := make([]Mapping, 0)
	err := yaml.Unmarshal(fieldRegistryFile, &fm)
	if err != nil {
		return nil, err
	}

	fr := make(map[string]Mapping, 0)
	for _, f := range fm {
		c := strings.TrimSpace(f.Campo)
		// To fix stuff like clienti[{0}].pec and transforming to clienti[].pec
		if strings.Index(f.DocumentMapping, "[{0}]") >= 0 {
			f.DocumentMapping = strings.ReplaceAll(f.DocumentMapping, "[{0}]", "[]")
		}
		fr[strings.ToLower(c)] = f
	}

	return fr, nil
}
