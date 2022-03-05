package tcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func (s *TCPServer) readKey(conn net.Conn) (*Frame, error) {
	frame := Frame{}
	err := binary.Read(conn, binary.BigEndian, &frame.Header)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, frame.Header.KeyLength)

	if _, err = io.ReadFull(conn, buf); err != nil {
		return nil, err
	}
	frame.Key = buf

	return &frame, nil
}

func (s *TCPServer) readKeyAndValue(conn net.Conn) (*Frame, error) {
	frame := Frame{}
	frame.Header = &FrameHeader{}
	err := binary.Read(conn, binary.BigEndian, frame.Header)
	if err != nil {
		return nil, err
	}
	// frame.Key = make([]byte, frame.Header.KeyLength)
	fmt.Println(frame.Header)
	keybuf := make([]byte, frame.Header.KeyLength)

	if _, err = io.ReadFull(conn, keybuf); err != nil {
		return nil, err
	}

	valuebuf := make([]byte, frame.Header.ValueLength)

	if _, err = io.ReadFull(conn, valuebuf); err != nil {
		return nil, err
	}

	frame.Key = keybuf
	frame.Value = valuebuf
	fmt.Println(frame, string(frame.Key), string(frame.Value))
	return &frame, nil
}
