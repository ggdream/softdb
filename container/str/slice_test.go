package str

import "testing"

func TestNewSlice(t *testing.T) {
	sliceStr := NewSlice()
	sliceStr.Set([]byte("我喜欢你"))
	sliceStr.Append([]byte("，哈哈"))
	println(sliceStr.StrLen())
	println(sliceStr.String())
	sliceStr.Set([]byte("123"))
}
