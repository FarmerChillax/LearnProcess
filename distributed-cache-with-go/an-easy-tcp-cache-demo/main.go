package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

type TestEncode struct {
	Key int64
}

func main() {

	// kvFrame := tcp.KeyValueFrame{
	// 	Key:   tcp.KeyFrame{Frame: tcp.Frame{Len: 1234}},
	// 	Value: tcp.ValueFrame{Frame: tcp.Frame{Len: 2333}},
	// }
	frame := TestEncode{Key: 233}
	buf := new(bytes.Buffer)

	fmt.Println(frame, buf)
	if err := binary.Write(buf, binary.BigEndian, &frame); err != nil {
		log.Println(err)
	}
	fmt.Println(buf.Bytes())
}

// func main() {
// 	buf := new(bytes.Buffer)
// 	var data = []interface{}{
// 		uint16(61374),
// 		int8(-54),
// 		uint8(254),
// 	}
// 	for _, v := range data {
// 		err := binary.Write(buf, binary.LittleEndian, v)
// 		if err != nil {
// 			fmt.Println("binary.Write failed:", err)
// 		}
// 	}
// 	fmt.Printf("%x", buf.Bytes())
// }
