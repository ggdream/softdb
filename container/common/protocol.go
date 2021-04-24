package common

import "time"


type CommonCmd interface {
	Keys() uint
	Del(key []byte) error
	// Select(db uint) error
	Exists(key []byte) bool
	// Type()
	Expire(key []byte, time time.Duration) error
	TTL(key []byte) time.Duration
	DBSize() uint
}
