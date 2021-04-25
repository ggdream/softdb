package server

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"net"
	"sync"
	"time"
	"unsafe"

	"github.com/ggdream/softdb/tools/pool"
)

type Server struct {
	mu       sync.Mutex
	closed   bool
	done     chan struct{}
	pool     *pool.Pool
	listener net.Listener
}

func NewServer() *Server {
	return &Server{
		done: make(chan struct{}),
		pool: pool.New(1 << 10),
	}
}

func (s *Server) Run(addr string) error {
	s.pool.Run()

	var err error
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Println("SoftDB is running, ready to accept connections.")

label:
	for {
		select {
		case <-s.done:
			break label
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				continue
			}
			s.pool.Add(func() {
				s.handleConn(conn)
			})
		}
	}

	return nil
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_ = conn.SetReadDeadline(time.Now().Add(time.Hour))
		bufReader := bufio.NewReader(conn)
		buf := make([]byte, 4)
		_, err := bufReader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("read cmd err: %+v\n", err)
			}
			break
		}
		size := binary.BigEndian.Uint32(buf[:4])
		if size <= 0 {
			log.Println("package size <=0 err")
			continue
		}

		data := make([]byte, size)
		if _, err = bufReader.Read(data); err != nil {
			log.Printf("read data err: %+v\n", err)
			break
		}
		reply := s.wrapReplyInfo(*(*string)(unsafe.Pointer(&data)))
		if _, err = conn.Write(reply); err != nil {
			log.Printf("write reply err: %+v\n", err)
		}
	}
}

func (s *Server) Stop() error {
	if s.closed {
		return errors.New("can't stop the server again")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	close(s.done)
	s.closed = true
	if err := s.listener.Close(); err != nil {
		return err
	}
	// TODO: close the db
	return nil
}

func (s *Server) wrapReplyInfo(data string) []byte {
	buf := make([]byte, len(data)+4)
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	_ = copy(buf[4:], data)
	return buf
}
