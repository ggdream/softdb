package softdb

import "unsafe"


// SoftDB
type SoftDB struct {
	db	map[string]interface{}
}

func New() *SoftDB {
	return &SoftDB{
		db: make(map[string]interface{}),
	}
}

func (s *SoftDB) slice2Str(value []byte) string {
	return *(*string)(unsafe.Pointer(&value))
}
