package str

type StringCmd interface {
	Set(value []byte)
	Get() []byte
	GetSet(value []byte) []byte
	Append(value []byte) []byte
	StrLen() int
	String() string
}
