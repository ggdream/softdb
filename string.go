package softdb

import "github.com/ggdream/softdb/tools/errno"

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


func (s *SoftDB) Set(key, value []byte) error {
	s.db[s.slice2Str(key)] = value
	return nil
}

func (s *SoftDB) Get(key []byte) ([]byte, error) {
	res, ok := s.db[s.slice2Str(key)]
	if !ok {
		return nil, errno.NotFoundTheKeyErr
	}
	return res.([]byte), nil
}

func (s *SoftDB) MSet(keysAndVals ...[]byte) error {
	return nil
}

func (s *SoftDB) MGet(keys ...[]byte) ([][]byte, error) {
	return nil, nil
}

func (s *SoftDB) SetNX(key, value []byte) error {
	return nil
}

func (s *SoftDB) GetSet(key, value []byte) ([]byte, error) {
	return nil, nil
}

func (s *SoftDB) Append(key, value []byte) error {
	return nil
}

func (s *SoftDB) StrLen(key []byte) (uint, error) {
	return 0, nil
}

func (s *SoftDB) Incr(key []byte) (int, error) {
	return 0, nil
}

func (s *SoftDB) Decr(key []byte) (int, error) {
	return 0, nil
}

func (s *SoftDB) IncrBy(key []byte, num int) (int, error) {
	return 0, nil
}

func (s *SoftDB) DecrBy(key []byte, num int) (int, error) {
	return 0, nil
}
