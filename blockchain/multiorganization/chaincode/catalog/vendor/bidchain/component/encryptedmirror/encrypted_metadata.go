package encryptedmirror

import (
	"bidchain/fabric/log"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
)

var (
	ErrFieldHashMapEmpty = errors.New("Invalid Empty HashMap")
	ErrSummaryMapInvalid = errors.New("Invalid empty SummaryMap")
	EmtpyStringHash256   []byte
)

func init() {
	sum256 := sha256.Sum256([]byte(""))
	EmtpyStringHash256 = sum256[:]
}

const (
	HASH_LENGTH = 32
)

type EncryptedDataMetadata struct {
	Title    string
	FieldMap map[string]interface{}
}

func (metadata *EncryptedDataMetadata) isValid(fieldHashMap map[string][]byte) (bool, error) {
	if len(fieldHashMap) == 0 {
		return false, ErrFieldHashMapEmpty;
	}
	validationMap := make(map[string][]byte)
	emptyHashCount := 0
	for key, value := range fieldHashMap {
		if bytes.Equal(value, EmtpyStringHash256) == true {
			emptyHashCount++
		}
		if _, exists := metadata.FieldMap[key]; !exists {
			return false, fmt.Errorf("title[%s] does not have field [%s]", metadata.Title, key)
		}
		//if _, ok := validationMap[key]; ok {
		//	return false, fmt.Errorf("title[%s] has duplicate field [%s]", metadata.Title, key)
		//}
		validationMap[key] = value
		if value == nil {
			msg := fmt.Sprintf("invalid SummaryInfoList, key=%s, hash is null", key)
			log.Warn(MODULE_NAME, msg)
			return false, errors.New(msg)
		}
		if len(value) != HASH_LENGTH {
			msg := fmt.Sprintf("invalid SummaryInfoList, key=%s, hash=[%v]'s length != 32", key, value)
			log.Warn(MODULE_NAME, msg)
			return false, errors.New(msg)
		}
	}
	if emptyHashCount == len(fieldHashMap) {
		return false, ErrSummaryMapInvalid
	}
	return true, nil
}
