package server

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	server := NewServer()
	// if err := server.Run(":8080"); err != nil {
	// 	panic(err)
	// }
	println(server)
}
