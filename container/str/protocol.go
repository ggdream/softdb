package str

type StringCmd interface {
	Set([]byte)
	Get() []byte
	Append([]byte) []byte
	Length() uint
	IsInt() bool
	String() string
}
