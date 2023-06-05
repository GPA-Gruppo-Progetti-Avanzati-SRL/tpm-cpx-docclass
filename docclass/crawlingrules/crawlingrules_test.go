package crawlingrules_test

import (
	_ "embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azbloblks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azstoragecfg"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/crawlingrules"
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

//go:embed crawling-rules.yml
var RulesConfigYaml []byte

func TestEmbeddedCrawlingRules(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfgs, err := crawlingrules.ReadYamlConfig(RulesConfigYaml)
	require.NoError(t, err)

	r, err := crawlingrules.NewRing(cfgs)
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
