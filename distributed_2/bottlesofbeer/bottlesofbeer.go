package main

import (
	"flag"
	"math/rand"
	"strconv"
	"time"
	//	"net/rpc"
	//	"fmt"
	//	"time"
	//	"net"
)

var nextAddr string

func main() {
	//	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	//	bottles := flag.Int("n",0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.
}

func getLyric(bottles int) string {
	// if sb calls this particular server's rpc, take its addr, decrement
	var lyric string
	time.Sleep(time.Duration(rand.Intn(bottles)) * time.Second)
	if bottles == 0 {
		lyric = "No more bottles of beer on the wall, no more bottles of beer. Take one down, pass it around..."
	} else {
		num := strconv.Itoa(bottles)
		lyric = num + " more bottles of beer on the wall, " + num + " more bottles of beer. Take one down, pass it around..."
	}
	return lyric
}
