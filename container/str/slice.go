// STR implemented by `go.native.[]byte`
package str

import (
	"strconv"
	"unicode/utf8"
	"unsafe"
)

type SliceStr struct {
	value []byte
	isInt bool
}

func NewSlice(value []byte) *SliceStr {
	s := &SliceStr{
		value: value,
	}
	s.judgeInt(value)
	return s
}

func (s *SliceStr) Set(value []byte) {
	s.value = value
	s.judgeInt(value)
}

func (s *SliceStr) Get() []byte {
	return s.value
}

func (s *SliceStr) Append(suffixValue []byte) []byte {
	s.value = append(s.value, suffixValue...)
	return s.value
}

func (s *SliceStr) Length() int {
	return utf8.RuneCountInString(s.String())
}

func (s *SliceStr) judgeInt(value []byte) {
	valStr := *(*string)(unsafe.Pointer(&value))
	_, err := strconv.Atoi(valStr)
	s.isInt = err == nil
}

func (s *SliceStr) IsInt() bool {
	return s.isInt
}

func (s *SliceStr) String() string {
	return *(*string)(unsafe.Pointer(&s.value))
}
