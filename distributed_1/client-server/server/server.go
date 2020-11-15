package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn *net.Conn) { // accept pointer to network connection
	reader := bufio.NewReader(*conn)
	for {
		msg, _ := reader.ReadString('\n') // terminated by newline character
		fmt.Printf(msg)
		fmt.Fprintln(*conn, "OK")
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8030") //open a port -look for tcp connection at port 8030; listen function creates servers
	for {
		conn, _ := ln.Accept() //accept connection
		go handleConnection(&conn)
	}
}
