package softdb


type ZSetCmd interface {
	ZAdd(key, member []byte, score float64) error
	ZRem(key, member []byte) (bool, error)
	ZScore(key, member []byte) float64
	ZCard(key []byte) uint
	ZRank(key, member []byte) (uint, error)
	ZRevRank(key, member []byte) (uint, error)
	ZIncrBy(key, member []byte, num float64) (float64, error)
	ZRange(key []byte, sIdx, eIdx int) ([][]byte, error)
	ZRevRange(key []byte, sIdx, eIdx int) ([][]byte, error)
	ZCount(key []byte) (uint, error)
	ZGetByRank(key []byte, rank int) (member []byte, score float64, err error)
	ZRevGetByRank(key []byte, rank int) (member []byte, score float64, err error)
	ZRangeByScore(key []byte, min, max float64) [][]byte
	ZRevRangeByScore(key []byte, min, max float64) [][]byte
}
