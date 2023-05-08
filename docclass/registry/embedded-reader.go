package registry

import (
	"embed"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"path"
	"tpm-cpx-docclass/docclass/model"
)

// The folder contains a number of .yml files each one for a different class
//
//go:embed defs/*
var docClassFS embed.FS
var docClassFSRootPath = "defs"

const YamlExtension = ".yml"

func ReadRegistryFromEmbeddedData() (int, error) {

	registry = make(map[string]model.DocClass, 0)
	dirEntries, err := docClassFS.ReadDir(docClassFSRootPath)
	if err != nil {
		return len(registry), err
	}

	for _, e := range dirEntries {

		fn := e.Name()

		if path.Ext(fn) != YamlExtension {
			continue
		}

		fileContent, err := docClassFS.ReadFile(path.Join(docClassFSRootPath, fn))
		if err != nil {
			return len(registry), err
		}

		dc, err := readDocClassYMLDefinition(fileContent)
		if err != nil {
			return len(registry), err
		}

		registry[fn[0:len(fn)-len(YamlExtension)]] = dc
	}

	return len(registry), nil
}

func readDocClassYMLDefinition(fileContent []byte) (model.DocClass, error) {

	var err error

	dc := struct {
		DocClass model.DocClass `yaml:"docClass"`
	}{}
	err = yaml.Unmarshal(fileContent, &dc)
	if err != nil {
		return model.DocClass{}, err
	}

	// do some computation on the loaded data.
	err = dc.DocClass.Finalize()
	if err != nil {
		log.Error().Err(err).Send()
	}

	return dc.DocClass, nil
}

func IsConfigured(dcId string) bool {
	_, ok := registry[dcId]
	return ok
}

func TraceRegistry() {
	log.Trace().Int("num-entries", len(registry)).Msg("document class registry")

	numFields := 0
	for n, entry := range registry {
		numFields += len(entry.Index)
		log.Trace().Int("num-indexes", len(entry.Index)).Str("docclass", n).Msg("document class info")
	}

	log.Trace().Int("total-num-indexes", numFields).Int("num-docclass", len(registry)).Msg("registry info")
}
