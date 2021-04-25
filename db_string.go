package softdb

import (
	"errors"
	"sync"

	"github.com/ggdream/softdb/container/str"
)


type StrIdx struct {
	mu sync.RWMutex
	indexes map[string]str.StringCmd
}

func newStrIdx() *StrIdx {
	return &StrIdx{
		indexes: make(map[string]str.StringCmd),
	}
}

func (db *SoftDB) Set(key, value []byte) error {
	if err := db.checkKeyValue(key, value); err != nil {
		return err
	}

	db.strIndex.mu.Lock()
	defer db.strIndex.mu.Unlock()

	if db.strIndex.indexes[db.slice2Str(key)] == nil {
		db.strIndex.indexes[db.slice2Str(key)] = str.NewSlice()
	}

	db.strIndex.indexes[db.slice2Str(key)].Set(value)
	return nil
}

func (db *SoftDB) Get(key []byte) []byte {
	if err := db.checkKeyValue(key); err != nil {
		return nil
	}

	db.strIndex.mu.RLock()
	defer db.strIndex.mu.RUnlock()

	return db.strIndex.indexes[db.slice2Str(key)].Get()
}

func (db *SoftDB) MSet(keysAndVals ...[]byte) error {
	for i := 0; i < len(keysAndVals)>>1; i+=2 {
		if err := db.Set(keysAndVals[i], keysAndVals[i+1]); err != nil {
			return err
		}
	}
	return nil
}

func (db *SoftDB) MGet(keys ...[]byte) [][]byte {
	var data [][]byte
	for _, v := range keys {
		res := db.Get(v)
		if res == nil {
			return nil
		}
		data = append(data, res)
	}
	return data
}

func (db *SoftDB) StrExist(key []byte) bool {
	_, ok := db.strIndex.indexes[db.slice2Str(key)]
	return ok
}

func (db *SoftDB) SetNX(key, value []byte) error {
	db.strIndex.mu.Lock()
	defer db.strIndex.mu.Unlock()

	if db.StrExist(key) {
		return nil
	}

	return db.Set(key, value)
}

func (db *SoftDB) GetSet(key, value []byte) ([]byte, error) {
	db.strIndex.mu.Lock()
	defer db.strIndex.mu.Unlock()

	if !db.StrExist(key) {
		return nil, errors.New("the key is empty")
	}

	return db.strIndex.indexes[db.slice2Str(key)].GetSet(value), nil
}

func (db *SoftDB) Append(key, value []byte) []byte {
	db.strIndex.mu.Lock()
	defer db.strIndex.mu.Unlock()

	return db.strIndex.indexes[db.slice2Str(key)].Append(value)
}

func (db *SoftDB) StrLen(key []byte) int {
	db.strIndex.mu.RLock()
	defer db.strIndex.mu.RUnlock()

	return db.strIndex.indexes[db.slice2Str(key)].StrLen()
}


func (db *SoftDB) Incr(key []byte) (int, error) {
	return 0, nil
}

func (db *SoftDB) Decr(key []byte) (int, error) {
	return 0, nil
}

func (db *SoftDB) IncrBy(key []byte, num int) (int, error) {
	return 0, nil
}

func (db *SoftDB) DecrBy(key []byte, num int) (int, error) {
	return 0, nil
}
