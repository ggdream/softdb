package softdb

import (
	"errors"
	"unsafe"
)



type (
	// SoftDB
	SoftDB struct {
		strIndex *StrIdx
		listIndex *ListIdx
	}
)

func New() *SoftDB {
	return &SoftDB{
		strIndex: newStrIdx(),
		listIndex: newListIdx(),
	}
}

func (db *SoftDB) slice2Str(value []byte) string {
	return *(*string)(unsafe.Pointer(&value))
}


func (db *SoftDB) checkKeyValue(key []byte, values ...[]byte) error {
	keySize := len(key)
	if keySize == 0 {
		return errors.New("length of key equals zero")
	}

	if len(key) > (1 << 10) {
		return errors.New("key length gt 1 << 10")
	}

	for _, v := range values {
		if len(v) > (1 << 13) {
			return errors.New("value length gt 1 << 13")
		}
	}

	return nil
}
