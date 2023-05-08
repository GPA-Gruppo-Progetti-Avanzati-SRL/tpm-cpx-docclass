package crawlingrules

import (
	"gopkg.in/yaml.v3"
)

func ReadYamlConfig(yamlData []byte) ([]Config, error) {
	var arr []Config
	err := yaml.Unmarshal(yamlData, &arr)

	if err != nil {
		return nil, err
	}

	return arr, nil
}
