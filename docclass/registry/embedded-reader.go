package registry

import (
	"embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/model"
	"path"
)

const YamlExtension = ".yml"

func ReadRegistryFromEmbeddedData(fs embed.FS, rootPath string) (int, error) {

	registry = make(map[string]model.DocClass, 0)
	dirEntries, err := fs.ReadDir(rootPath)
	if err != nil {
		return len(registry), err
	}

	for _, e := range dirEntries {

		fn := e.Name()

		if path.Ext(fn) != YamlExtension {
			continue
		}

		fileContent, err := fs.ReadFile(path.Join(rootPath, fn))
		if err != nil {
			return len(registry), err
		}

		dc, err := model.ReadDocClassYMLDefinition(fn, fileContent)
		if err != nil {
			return len(registry), err
		}

		registry[fn[0:len(fn)-len(YamlExtension)]] = dc
	}

	return len(registry), nil
}

//func readDocClassYMLDefinition(fn string, fileContent []byte) (model.DocClass, error) {
//
//	const semLogContext = "doc-class-registry::read-doc-class-from-yaml"
//	log.Info().Str("fn", fn).Msg(semLogContext)
//	var err error
//
//	dc := struct {
//		DocClass model.DocClass `yaml:"docClass"`
//	}{}
//	err = yaml.Unmarshal(fileContent, &dc)
//	if err != nil {
//		return model.DocClass{}, err
//	}
//
//	// do some computation on the loaded data.
//	err = dc.DocClass.Finalize()
//	if err != nil {
//		log.Error().Err(err).Msg(semLogContext)
//	}
//
//	return dc.DocClass, nil
//}
