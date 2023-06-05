package registry_test

import (
	"embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azbloblks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azstoragecfg"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/fielddictionary"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/registry"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// The folder contains a number of .yml files each one for a different class
//
//go:embed defs/*
var EmbeddedDocClassFS embed.FS
var EmbeddedDocClassFSRootPath = "defs"

//go:embed embedded-dictionary.yml
var EmbeddedDictionary []byte

const (
	TargetContainer = "cpx-docclass-container"

	AZCommonBlobAccountNameEnvVarName = "AZCOMMON_BLOB_ACCOUNTNAME"
	AZCommonBlobAccountKeyEnvVarName  = "AZCOMMON_BLOB_ACCTKEY"
)

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
