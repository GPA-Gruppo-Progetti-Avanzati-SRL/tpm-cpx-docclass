package registry

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azbloblks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-cpx-docclass/docclass/model"
	"github.com/rs/zerolog/log"
	"regexp"
)

var RegExpFilter = regexp.MustCompile("^pt(.)*yml$")

func ReadRegistryFromBlobStorage(storageName, cntName string) (int, error) {

	const semLogContext = "doc-class-registry::read-from-blob-storage"

	registry = make(map[string]model.DocClass)

	lks, err := azbloblks.GetLinkedService(storageName)
	if err != nil {
		return len(registry), err
	}

	bi, err := lks.ListBlobs(cntName, 5000)
	if err != nil {
		return len(registry), err
	}

	for _, e := range bi {

		fn := e.BlobName

		if !RegExpFilter.Match([]byte(fn)) {
			log.Info().Str("blob-name", fn).Msg(semLogContext + " skipping blob")
			continue
		}

		fileContent, err := lks.DownloadToBuffer(cntName, fn)
		if err != nil {
			return len(registry), err
		}

		dc, err := model.ReadDocClassYMLDefinition(fn, fileContent.Body)
		if err != nil {
			return len(registry), err
		}

		registry[fn[0:len(fn)-len(YamlExtension)]] = dc
	}

	return len(registry), nil
}
