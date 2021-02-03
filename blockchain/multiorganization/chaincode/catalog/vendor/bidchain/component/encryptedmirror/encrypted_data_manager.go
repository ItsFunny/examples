package encryptedmirror

import (
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/log"
	"fmt"
)


var (
	MODULE_NAME = bsmodule.ENCRYPTED_DATA_MANAGER
)

type IEncryptedDataManager interface {
	AddMetadata(metadata *EncryptedDataMetadata)
	Validate(title string, fieldHash map[string][]byte) error
}

type DefaultEncryptedDataManager struct {
	IEncryptedDataManager
	MetadataMap map[string]*EncryptedDataMetadata //key title name, value map
}

func NewEncryptedDataManager() IEncryptedDataManager {
	encryptedDataMgr := DefaultEncryptedDataManager{}
	encryptedDataMgr.MetadataMap = make(map[string]*EncryptedDataMetadata)
	return &encryptedDataMgr
}

func (m *DefaultEncryptedDataManager) AddMetadata(metadata *EncryptedDataMetadata) {
	if _, exists:=m.MetadataMap[metadata.Title]; exists {
		m.MetadataMap[metadata.Title] = metadata
		log.Infof(MODULE_NAME, "AddMetadata found duplicate title : %s", metadata.Title)
	} else {
		m.MetadataMap[metadata.Title] = metadata
	}
}

func (m *DefaultEncryptedDataManager) Validate(title string, fieldHash map[string][]byte) error {
	if value, exists:=m.MetadataMap[title]; exists {
		_, err := value.isValid(fieldHash)
		return err
	} else {
		log.Warnf(MODULE_NAME, "Unknown title [%s] in DefaultEncryptedDataManager", title)
		return fmt.Errorf("Unknown title [%s] in DefaultEncryptedDataManager", title)
	}
}
