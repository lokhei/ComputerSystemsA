package main

import (
	//	"errors"
	"flag"
	//	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"secretstrings/stubs"
	"time"
)

/** Super-Secret `reversing a string' method we can't allow clients to see. **/
func ReverseString(s string, i int) string {
	time.Sleep(time.Duration(rand.Intn(i)) * time.Second)
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type SecretStringOperations struct{}

/**ReverseString **/
func (s *SecretStringOperations) Reverse(req *stubs.Request, res *stubs.Response) (err error) {
	res.Message = ReverseString(req.Message, 10)
	return
}

/**ReverseString - premium **/
func (s *SecretStringOperations) FastReverse(req *stubs.Request, res *stubs.Response) (err error) {
	res.Message = ReverseString(req.Message, 2) //decrease random sleep time
	return
}

func main() {
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	rpc.Register(&SecretStringOperations{})
	listener, _ := net.Listen("tcp", ":"+*pAddr)
	rpc.Accept(listener)

}
