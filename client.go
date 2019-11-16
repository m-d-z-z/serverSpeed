package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func runClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 64 * 1000)
	var totalTrafficLocal int64 = 0
	for {
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
			return
		}
		totalTrafficLocal += int64(size)
		select {
		case <-totalTraffic:
		default:
			log.Println("skip")
		}
		totalTraffic <- totalTrafficLocal
	}
}