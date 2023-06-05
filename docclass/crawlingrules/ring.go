package crawlingrules

import (
	_ "embed"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/registry"
	"github.com/rs/zerolog/log"
)

func NewRing(cfg []Config) (RuleRing, error) {

	const semLogContext = "doc-class-yaml::new-ring"

	ring := RuleRing{cursor: -1}

	p := make([]Rule, 0)
	checkMap := make(map[string]struct{})

	numDisabled := 0
	for _, pcfg := range cfg {
		if _, ok := checkMap[pcfg.Name]; ok {
			log.Warn().Str("name", pcfg.Name).Msg(semLogContext + " duplicate reference to doc-class... skipping")
			continue
		}

		checkMap[pcfg.Name] = struct{}{}

		if _, ok := registry.GetDocClass(pcfg.Name); !ok {
			log.Warn().Str("name", pcfg.Name).Msg(semLogContext + " doc-class referenced in crawler rules but not found in registry... skipping")
			continue
		}

		if pcfg.Disabled {
			numDisabled++
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

	log.Info().Int("num-disabled", numDisabled).Int("ring-size", len(p)).Int("doc-class-registry-size", registry.Size()).Msg(semLogContext)
	if len(p) == 0 {
		return ring, errors.New("no crawling rules defined or enabled")
	}

	ring.rules = p
	return ring, nil
}
