package main

import (
	"encoding/binary"
	"net"
)

const (
	TYPE_STRING uint8 = iota + 1
	TYPE_BINAIRE
	TYPE_INT
)

func main() {
	server, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 8000})
	for {
		conn, _ := server.AcceptTCP()
		go manage(conn)
	}
	if err != nil {
		panic(err)
	}
}

func manage(conn *net.TCPConn) {
	binary.Write(conn, binary.BigEndian, TYPE_STRING)
	var payload = []byte("blalbgskqwelwle \n beqwkd-oqkeodk \n wqjei0dhej9uewqjd \n")
	binary.Write(conn, binary.BigEndian, uint32(len(payload)))
	conn.Write(payload)
}
