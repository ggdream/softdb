package set


type SetCmd interface {
	SAdd(key []byte, members ...[]byte) (uint, error)
	SPop(key []byte, count int) ([][]byte, error)
	SRem(key []byte, members ...[]byte) (uint, error)
	SMove(from, to, member []byte) error
	SIsMember(key, member []byte) bool
	Scard(key []byte) uint
	SMembers(key []byte) [][]byte
	SRandMembers(key []byte, count int) [][]byte
	SInter(keys ...[]byte) [][]byte
	SUnion(keys ...[]byte) [][]byte
	SDiff(key1, key2 []byte) [][]byte
}
