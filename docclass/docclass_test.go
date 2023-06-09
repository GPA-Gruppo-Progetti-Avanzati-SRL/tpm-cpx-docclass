package docclass_test

import (
	"embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azbloblks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azstoragecfg"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/crawlingrules"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/fielddictionary"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/registry"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	TargetContainer = "cpx-docclass-container"

	AZCommonBlobAccountNameEnvVarName = "AZCOMMON_BLOB_ACCOUNTNAME"
	AZCommonBlobAccountKeyEnvVarName  = "AZCOMMON_BLOB_ACCTKEY"
)

// The folder contains a number of .yml files each one for a different class
//
//go:embed defs/*
var EmbeddedDocClassFS embed.FS
var EmbeddedDocClassFSRootPath = "defs"

//go:embed embedded-dictionary.yml
var EmbeddedDictionary []byte

//go:embed crawling-rules.yml
var RulesConfigYaml []byte

func TestEmbeddedCrawlingRules(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_, err := crawlingrules.ReadYamlConfig(RulesConfigYaml)
	require.NoError(t, err)

	r, err := crawlingrules.NewRing(true)
	require.NoError(t, err)

	for {
		rule, ok := r.Next()
		if !ok {
			return
		}

		t.Log(rule)
		rule.Crawl()
	}
}

func TestBlobStorageCrawlingRules(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	stgConfig := azstoragecfg.Config{
		Name:       "default",
		Account:    os.Getenv(AZCommonBlobAccountNameEnvVarName),
		AccountKey: os.Getenv(AZCommonBlobAccountKeyEnvVarName),
		AuthMode:   azstoragecfg.AuthModeAccountKey,
	}

	require.NotEmpty(t, stgConfig.Account, "blob storage account-name not set.... use env var "+AZCommonBlobAccountNameEnvVarName)
	require.NotEmpty(t, stgConfig.AccountKey, "blob storage account-key not set.... use env var "+AZCommonBlobAccountKeyEnvVarName)

	_, err := azbloblks.Initialize([]azstoragecfg.Config{stgConfig})
	require.NoError(t, err)

	rs, err := crawlingrules.ReadRulesFromBlobStorage("default", TargetContainer)
	require.NoError(t, err)

	t.Log("num rules: ", len(rs))
}

func TestEmbeddedRegistry(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	_, err := fielddictionary.ReadDictionaryFromYamlData(EmbeddedDictionary)
	require.NoError(t, err)

	_, err = registry.ReadRegistryFromEmbeddedData(EmbeddedDocClassFS, EmbeddedDocClassFSRootPath)
	require.NoError(t, err)

	registry.TraceRegistry()
}

func TestBlobStorageRegistry(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	stgConfig := azstoragecfg.Config{
		Name:       "default",
		Account:    os.Getenv(AZCommonBlobAccountNameEnvVarName),
		AccountKey: os.Getenv(AZCommonBlobAccountKeyEnvVarName),
		AuthMode:   azstoragecfg.AuthModeAccountKey,
	}

	require.NotEmpty(t, stgConfig.Account, "blob storage account-name not set.... use env var "+AZCommonBlobAccountNameEnvVarName)
	require.NotEmpty(t, stgConfig.AccountKey, "blob storage account-key not set.... use env var "+AZCommonBlobAccountKeyEnvVarName)

	_, err := azbloblks.Initialize([]azstoragecfg.Config{stgConfig})
	require.NoError(t, err)

	_, err = fielddictionary.ReadDictionaryFromBlobStorage("default", TargetContainer)
	require.NoError(t, err)

	_, err = registry.ReadRegistryFromBlobStorage("default", TargetContainer)
	require.NoError(t, err)

	registry.TraceRegistry()
}

func TestDocClassEmbeddedConfig(t *testing.T) {
	err := docclass.InitDocumentClassRegistry(docclass.DocumentClassConfig{
		Mode:           "embedded",
		FS:             EmbeddedDocClassFS,
		FSRootPath:     EmbeddedDocClassFSRootPath,
		CrawlingRules:  RulesConfigYaml,
		FileDictionary: EmbeddedDictionary,
	})

	require.NoError(t, err)
}

func TestDocClassBlobStorageConfig(t *testing.T) {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	stgConfig := azstoragecfg.Config{
		Name:       "default",
		Account:    os.Getenv(AZCommonBlobAccountNameEnvVarName),
		AccountKey: os.Getenv(AZCommonBlobAccountKeyEnvVarName),
		AuthMode:   azstoragecfg.AuthModeAccountKey,
	}

	require.NotEmpty(t, stgConfig.Account, "blob storage account-name not set.... use env var "+AZCommonBlobAccountNameEnvVarName)
	require.NotEmpty(t, stgConfig.AccountKey, "blob storage account-key not set.... use env var "+AZCommonBlobAccountKeyEnvVarName)

	_, err := azbloblks.Initialize([]azstoragecfg.Config{stgConfig})
	require.NoError(t, err)

	err = docclass.InitDocumentClassRegistry(docclass.DocumentClassConfig{
		Mode:        "blob",
		StorageName: "default",
		Container:   TargetContainer,
	})

	require.NoError(t, err)
}
