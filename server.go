package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func runServer(port string)  {
	fmt.Println("listen on " + port)
	fmt.Println("Please run this command in client machine: \n./serverSpeed client " + getMyIP() + port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go serverHandler(conn)
	}
}

func getMyIP() string {
	res, err := http.Get("https://api.ipify.org")
	if err != nil {
		panic(err)
	}
	ip, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(ip)
}

func serverHandler(conn net.Conn) {
	buf := make([]byte, 64 * 1000)
	packID := 0
	for {
		if packID%7 == 0 {
			rand.Read(buf)
		}
		packID++
		log.Printf("id: %d sending to " + conn.RemoteAddr().String() + "\n", packID)
		_, err := conn.Write(buf)
		if err != nil {
			log.Println(err)
			return
		}
	}
}