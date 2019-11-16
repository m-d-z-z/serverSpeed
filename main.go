package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ServerSpeed v0.0.4 by @Undefiner (mdzz@mdzz.moe)\n机场技术支持、开发可以找我")
	if len(os.Args) != 3 {
		help()
		return
	}
	mode := os.Args[1]
	addr := os.Args[2]
	switch mode {
	case "server":
		runServer(addr)
	case "client":
		go runClient(addr)
		trafficPrinter()
	default:
		help()
	}
}

func help() {
	fmt.Println("serverspeed [server/client] [address/port]")
	return
}