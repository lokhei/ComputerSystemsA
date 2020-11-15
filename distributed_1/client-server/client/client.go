package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}
func main() {
	//build new reader that reads from standard input
	stdin := bufio.NewReader(os.Stdin)
	conn, _ := net.Dial("tcp", "127.0.0.1:8030") // Dial function connects to a server
	// msgPtr := flag.String("msg", "Default message", "Message to send.") // get message pointer fom flag package
	// // refer to it as msg, set default msg, give string explainign what this variable is for
	// flag.Parse()                                     // to parse all the flags
	for {
		fmt.Println("Enter Text: ")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, text) // write message to connection
		read(&conn)
	}
}
