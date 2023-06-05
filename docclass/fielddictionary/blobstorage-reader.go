package fielddictionary

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-az-common/storage/azbloblks"
	"github.com/rs/zerolog/log"
)

func ReadDictionaryFromBlobStorage(storageName, cntName string) (FieldDictionary, error) {

	const semLogContext = "field-dictionary::read-from-blob-storage"

	lks, err := azbloblks.GetLinkedService(storageName)
	if err != nil {
		return nil, err
	}

	bi, err := lks.DownloadToBuffer(cntName, "field-registry.yml")
	if err != nil {
		return nil, err
	}

	log.Info().Msg(semLogContext)
	return ReadDictionaryFromYamlData(bi.Body)
}
