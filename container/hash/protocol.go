package hash


type HashCmd interface {
	HSet(key, field, value []byte) (uint, error)
	HGet(key, field []byte) ([]byte, error)
	HMSet(key []byte, fieldsAndVals ...[]byte) (uint, error)
	HMGet(key []byte, fields ...[]byte) ([][]byte, error)
	HSetNX(key, field []byte) error
	HExists(key, field []byte) (bool, error)
	HKeys(key []byte) ([][]byte, error)
	HVals(key []byte) ([][]byte, error)
	HDel(key, field []byte) (bool, error)
	HIncrBy(key, field []byte, num int) (int, error)
}
