package tcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func Client(server, port, option, key, value string) error {
	frame := Frame{}
	frame.init([]byte(key), []byte(value))

	clientAddr := net.JoinHostPort(server, port)
	client, err := net.Dial("tcp", clientAddr)
	defer client.Close()
	if err != nil {
		return err
	}
	switch option {
	case "get":
		method := []byte("G")
		_, err = client.Write(method)
		if err != nil {
			return err
		}
		// send request
		if err = sendRequest(&frame, client); err != nil {
			return err
		}

		// recv response
		frame, err := recvResponse(client)
		if err != nil {
			return err
		}
		fmt.Println(frame)
	case "set":
		method := []byte("S")
		if _, err = client.Write(method); err != nil {
			return err
		}

		if err = sendRequest(&frame, client); err != nil {
			return err
		}
		fmt.Println("发送请求完成")

		frame, err := recvResponse(client)
		if err != nil {
			return err
		}
		fmt.Println("switch:", frame)
	case "del":
	}
	return nil
}

func sendRequest(frame *Frame, conn net.Conn) (err error) {
	if err = binary.Write(conn, binary.BigEndian, frame.Header); err != nil {
		return err
	}

	if _, err = conn.Write(frame.Key); err != nil {
		return err
	}

	_, err = conn.Write(frame.Value)
	return err
}

func recvResponse(conn net.Conn) (*Frame, error) {
	frame := Frame{}
	frame.Header = &FrameHeader{}
	err := binary.Read(conn, binary.BigEndian, frame.Header)
	if err != nil {
		return nil, err
	}

	// recv value
	buf := make([]byte, frame.Header.ValueLength)
	if _, err = io.ReadFull(conn, buf); err != nil {
		return nil, err
	}

	fmt.Println("recv:", frame, string(frame.Key))
	return &frame, err

}
