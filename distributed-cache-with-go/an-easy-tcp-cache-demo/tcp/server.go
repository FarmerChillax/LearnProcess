package tcp

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"tcp-cache-service/cache"
)

type TCPServer struct {
	cache.Cache
}

func (s *TCPServer) Listen() {
	listen, err := net.Listen("tcp", ":2333")
	if err != nil {
		panic(err)
	}
	for {
		con, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go s.process(con)
	}
}

func (s *TCPServer) process(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		op, err := r.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("close connection due to err:", err)
			}
			return
		}
		fmt.Println(string(op))
		fmt.Println(s.Cache.GetStat())
		switch op {
		case 'S':
			err = s.set(conn)
			fmt.Println("发送响应完成, err:", err)
		case 'G':
			err = s.get(conn)
		case 'D':
			err = s.del(conn)
		default:
			log.Println()
		}
		if err != nil {
			log.Println("close connection due to error:", err)
			return
		}
	}
}

// get method
func (s *TCPServer) get(conn net.Conn) error {
	frame, err := s.readKey(conn)
	if err != nil {
		return err
	}

	value, err := s.Cache.Get(string(frame.Key))
	if err != nil {
		return err
	}

	frame.Value = value

	return sendResponse(frame, err, conn)
}

// set method
func (s *TCPServer) set(conn net.Conn) error {
	frame, err := s.readKeyAndValue(conn)
	if err != nil {
		return err
	}

	return sendResponse(frame, s.Cache.Set(string(frame.Key), frame.Value), conn)
}

// del method
func (s *TCPServer) del(conn net.Conn) error {
	frame, err := s.readKey(conn)
	if err != nil {
		return err
	}

	return sendResponse(frame, s.Cache.Del(string(frame.Key)), conn)
}

func New(c cache.Cache) *TCPServer {
	return &TCPServer{c}
}

func sendResponse(frame *Frame, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("%-d %s", len(errString), errString)
		_, err = conn.Write([]byte(tmp))
		fmt.Println("出错了")
		return err
	}
	if err = binary.Write(conn, binary.BigEndian, frame.Header); err != nil {
		fmt.Println("发送header出错")
		return err
	}

	_, err = conn.Write(frame.Value)
	return err
}
