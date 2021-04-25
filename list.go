package softdb


type (
	InsertType	int

	ListCmd interface {
		Lpush(key []byte, values ...[]byte) (uint, error)
		RPush(key []byte, values ...[]byte) (uint, error)
		LPop(key []byte) ([]byte, error)
		RPop(key []byte) ([]byte, error)
		LRange(key []byte, sIdx, eIdx int) ([][]byte, error)
		LIndex(key []byte, idx int) ([]byte, error)
		LSet(key, value []byte, idx int) error
		LRem(key, value []byte, count int) (uint, error)
		LInsert(key []byte, insertType InsertType, pivot, value []byte) (uint, error)
	}
)

const (
	InsertTypeBefore InsertType = iota
	InsertTypeAfter
)
