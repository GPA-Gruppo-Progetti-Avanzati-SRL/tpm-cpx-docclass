package crawlingrules_test

import (
	_ "embed"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"tpm-cpx-docclass/docclass/crawlingrules"
)

func TestCrawlingRules(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	cfgs, err := crawlingrules.ReadYamlConfig(crawlingrules.RulesConfigYaml)
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
