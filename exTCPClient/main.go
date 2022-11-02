package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{Port: 8000})
	if err != nil {
		panic(err)
	}

	var t uint8
	var l uint32
	binary.Read(conn, binary.BigEndian, &t)
	if t == 1 {
		binary.Read(conn, binary.BigEndian, &l)
		var v = make([]byte, l)
		conn.Read(v)
		fmt.Print(string(v))
	}

}
