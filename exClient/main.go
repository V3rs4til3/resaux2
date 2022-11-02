package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	dialer, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	for {
		msg, err := bufio.NewReader(dialer).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("-> : ", msg)
		time.Sleep(1 * time.Second)
	}
}
