package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Hello world")
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	handleConn(conn)
}

func handleConn(conn net.Conn) {
	fmt.Fprintln(conn, "some message")
}
