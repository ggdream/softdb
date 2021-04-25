package softdb

import (
	"sync"

	"github.com/ggdream/softdb/container/list"
)


type ListIdx struct {
	mu sync.RWMutex
	indexes	map[string]list.ListCmd
}

func newListIdx() *ListIdx {
	return &ListIdx{
		indexes: make(map[string]list.ListCmd),
	}
}

func (db *SoftDB) LPush(key []byte, values ...[]byte) (int, error) {
	if err := db.checkKeyValue(key, values...); err != nil {
		return 0, err
	}

	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	if db.listIndex.indexes[db.slice2Str(key)] == nil {
		db.listIndex.indexes[db.slice2Str(key)] = list.NewDoubleLinkedList()
	}
	db.listIndex.indexes[db.slice2Str(key)].LPush(values...)
	return db.listIndex.indexes[db.slice2Str(key)].LLen(), nil
}

func (db *SoftDB) RPush(key []byte, values ...[]byte) (int, error) {
	if err := db.checkKeyValue(key, values...); err != nil {
		return 0, err
	}

	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	if db.listIndex.indexes[db.slice2Str(key)] == nil {
		db.listIndex.indexes[db.slice2Str(key)] = list.NewDoubleLinkedList()
	}
	db.listIndex.indexes[db.slice2Str(key)].RPush(values...)
	return db.listIndex.indexes[db.slice2Str(key)].LLen(), nil
}

func (db *SoftDB) LPop(key []byte) ([]byte, error) {
	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	return db.listIndex.indexes[db.slice2Str(key)].LPop()
}

func (db *SoftDB) RPop(key []byte) ([]byte, error) {
	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	return db.listIndex.indexes[db.slice2Str(key)].RPop()
}

func (db *SoftDB) LIndex(key []byte, idx int) ([]byte, error) {
	db.listIndex.mu.RLock()
	defer db.listIndex.mu.RUnlock()

	return db.listIndex.indexes[db.slice2Str(key)].LIndex(idx)
}

func (db *SoftDB) LRem(key, value []byte, count int) int {
	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	return db.listIndex.indexes[db.slice2Str(key)].LRem(value, count)
}

func (db *SoftDB) LInsert(key []byte, insertType list.InsertType, pivot, value []byte) (int, error) {
	if err := db.checkKeyValue(key, value); err != nil {
		return 0, err
	}

	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	return db.listIndex.indexes[db.slice2Str(key)].LInsert(insertType, pivot, value), nil
}

func (db *SoftDB) LSet(key []byte, idx int, value []byte) error {
	db.listIndex.mu.Lock()
	defer db.listIndex.mu.Unlock()

	return db.listIndex.indexes[db.slice2Str(key)].LSet(idx, value)
}

func (db *SoftDB) LRange(key []byte, sIdx, eIdx int) ([][]byte, error) {
	if err := db.checkKeyValue(key); err != nil {
		return nil, err
	}

	db.listIndex.mu.RLock()
	defer db.listIndex.mu.RUnlock()

	return db.listIndex.indexes[db.slice2Str(key)].LRange(sIdx, eIdx)
}

func (db *SoftDB) LLen(key []byte) int {
	db.listIndex.mu.RLock()
	defer db.listIndex.mu.RUnlock()

	return db.listIndex.indexes[db.slice2Str(key)].LLen()
}
