package crawlingrules

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azbloblks"
	"github.com/rs/zerolog/log"
)

func ReadRulesFromBlobStorage(storageName, cntName string) ([]Config, error) {

	const semLogContext = "crawling-rules::read-from-blob-storage"

	lks, err := azbloblks.GetLinkedService(storageName)
	if err != nil {
		return nil, err
	}

	bi, err := lks.DownloadToBuffer(cntName, "crawling-rules.yml")
	if err != nil {
		return nil, err
	}

	log.Info().Msg(semLogContext)
	return ReadYamlConfig(bi.Body)
}
