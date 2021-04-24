package str


type StringCmd interface {
	Set(key, value []byte) error
	Get(key []byte) ([]byte, error)
	MSet(keysAndVals ...[]byte) error
	MGet(keys ...[]byte) ([][]byte, error)
	SetNX(key, value []byte) error
	GetSet(key, value []byte) ([]byte, error)
	Append(key, value []byte) error
	StrLen(key []byte) (uint, error)
	Incr(key []byte) (int, error)
	Decr(key []byte) (int, error)
	IncrBy(key []byte, num int) (int, error)
	DecrBy(key []byte, num int) (int, error)
}
