package list


type (
	InsertType	int

	ListCmd interface {
		LPush(values ...[]byte) int
		RPush(values ...[]byte) int
		LPop() ([]byte, error)
		RPop() ([]byte, error)
		LRange(sIdx, eIdx int) ([][]byte, error)
		LIndex(idx int) ([]byte, error)
		LInsert(insertType InsertType, pivot, value []byte) int
		LLen() int
		LSet(idx int, value []byte) error
		LRem(value []byte, count int) int
	}
)

const (
	Before InsertType = iota
	After
)
