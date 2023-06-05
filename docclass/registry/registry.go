package registry

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/model"
	"github.com/rs/zerolog/log"
)

var registry map[string]model.DocClass

func GetRegistry() map[string]model.DocClass {
	return registry
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
