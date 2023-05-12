package registry

import (
	"embed"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/model"
)

// The folder contains a number of .yml files each one for a different class
//
//go:embed defs/*
var EmbeddedDocClassFS embed.FS
var EmbeddedDocClassFSRootPath = "defs"

var registry map[string]model.DocClass

func GetRegistry() map[string]model.DocClass {
	return registry
}
