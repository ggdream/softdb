package softdb

import "testing"

func TestStrIdx(t *testing.T) {
	db := New()
	if err := db.Set([]byte("name"), []byte("wang")); err != nil {
		panic(err)
	}
	println(db.slice2Str(db.Get([]byte("name"))))
}
