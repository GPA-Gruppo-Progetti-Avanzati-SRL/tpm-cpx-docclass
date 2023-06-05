package docclass

import (
	"embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/crawlingrules"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/fielddictionary"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/registry"
	"github.com/rs/zerolog/log"
)

type DocumentClassConfig struct {
	Mode           string   `mapstructure:"mode" yaml:"mode" json:"mode"`
	StorageName    string   `mapstructure:"storage-name,omitempty" yaml:"storage-name,omitempty" json:"storage-name,omitempty"`
	Container      string   `mapstructure:"container,omitempty" yaml:"container,omitempty" json:"container,omitempty"`
	CrawlingRules  []byte   `mapstructure:"-" yaml:"-" json:"-"`
	FileDictionary []byte   `mapstructure:"-" yaml:"-" json:"-"`
	FS             embed.FS `mapstructure:"-" yaml:"-" json:"-"`
	FSRootPath     string   `mapstructure:"-" yaml:"-" json:"-"`
}

func InitDocumentClassRegistry(cfg DocumentClassConfig) error {

	const semLogContext = "doc-class::init-registry"

	var err error
	if cfg.Mode == "blob" {
		err = initFromBlobStorage(cfg.StorageName, cfg.Container)
	} else {
		err = initFromEmbeddedConfig(&cfg)
	}

	if err == nil {
		registry.TraceRegistry()
		log.Info().Int("num-classes", registry.Size()).Msg(semLogContext + " document class registry loaded")
	}

	return err
}

func initFromEmbeddedConfig(cfg *DocumentClassConfig) error {
	_, err := crawlingrules.ReadYamlConfig(cfg.CrawlingRules)
	if err != nil {
		return err
	}

	_, err = fielddictionary.ReadDictionaryFromYamlData(cfg.FileDictionary)
	if err != nil {
		return err
	}

	_, err = registry.ReadRegistryFromEmbeddedData(cfg.FS, cfg.FSRootPath)
	return err
}

func initFromBlobStorage(name string, container string) error {
	_, err := crawlingrules.ReadRulesFromBlobStorage(name, container)
	if err != nil {
		return err
	}

	_, err = fielddictionary.ReadDictionaryFromBlobStorage(name, container)
	if err != nil {
		return err
	}

	_, err = registry.ReadRegistryFromBlobStorage(name, container)
	return err
}
