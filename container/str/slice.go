// STR implemented by `go.native.[]byte`
package str

import (
	"unicode/utf8"
	"unsafe"
)

type SliceStr struct {
	value []byte
}

func NewSlice() *SliceStr {
	return &SliceStr{}
}

func (s *SliceStr) Set(value []byte) {
	s.value = value
}

func (s *SliceStr) Get() []byte {
	return s.value
}

func (s *SliceStr) GetSet(value []byte) []byte {
	temp := s.value
	s.value = value
	return temp
}

func (s *SliceStr) Append(value []byte) []byte {
	s.value = append(s.value, value...)
	return s.value
}

func (s *SliceStr) StrLen() int {
	return utf8.RuneCountInString(s.String())
}


func (s *SliceStr) String() string {
	return *(*string)(unsafe.Pointer(&s.value))
}
