package crawlingrules

import (
	_ "embed"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/rs/zerolog/log"
)

//go:embed crawling-rules.yml
var RulesConfigYaml []byte

func NewRing(cfg []Config) (RuleRing, error) {

	const semLogContext = "doc-class-yaml::new-ring"

	ring := RuleRing{cursor: -1}

	p := make([]Rule, 0)
	for _, pcfg := range cfg {
		if pcfg.Disabled {
			log.Info().Str("class_id", pcfg.Name).Msg(semLogContext + " rule disabled")
			continue
		}

		tr, err := util.NewTimeOfDayRangesFromString(pcfg.TimeRange)
		if err != nil {
			log.Error().Err(err).Str("time-range", pcfg.TimeRange).Msg(semLogContext + " time range check disabled")
		}

		dcr := Rule{cfg: pcfg, validIn: tr}
		p = append(p, dcr)
	}

	if len(p) == 0 {
		return ring, errors.New("no crawling rules defined or enabled")
	}

	log.Info().Int("num-rules", len(p)).Msg(semLogContext + " rules configured")
	ring.rules = p
	return ring, nil
}
