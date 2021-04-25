package str

import "testing"

func TestNewSlice(t *testing.T) {
	sliceStr := NewSlice([]byte("我喜欢你"))
	sliceStr.Append([]byte("，哈哈"))
	println(sliceStr.Length())
	println(sliceStr.IsInt())
	println(sliceStr.String())
	sliceStr.Set([]byte("123"))
	println(sliceStr.IsInt())
}
