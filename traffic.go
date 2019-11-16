package main

import (
	"fmt"
	"time"
)

var totalTraffic = make(chan int64, 1)
var (
	lastCheckTraffic int64
	lastCheckTime int64
)

func trafficPrinter() {
	fmt.Println("listening")
    var totalTrafficLocal int64 = 0
	for {
		totalTrafficLocal = <- totalTraffic
		thisCheckTime := time.Now().Unix()
		speed := (totalTrafficLocal - lastCheckTraffic) / (thisCheckTime - lastCheckTime)
		lastCheckTraffic = totalTrafficLocal
		lastCheckTime = thisCheckTime
		fmt.Println(formatSpeedTraffic(speed) + "/s")
		time.Sleep(time.Second)
	}
}

func formatSpeedTraffic(traffic int64) string {
	if traffic < 1024 {
		return fmt.Sprintf("%.2f B", float64(traffic))
	}
	if traffic <1024 * 1024 {
		return fmt.Sprintf("%.2f KB", float64(traffic) / 1024)
	}
	if traffic <1024 * 1024 * 1024 {
		return fmt.Sprintf("%.2f MB", float64(traffic) / 1024 / 1024)
	}
	if traffic <1024 * 1024 * 1024 * 1024 {
		return fmt.Sprintf("%.2f GB", float64(traffic) / 1024 / 1024 / 1024)
	}
	return fmt.Sprintf("%.2f TB", float64(traffic) / 1024 / 1024 / 1024 / 1024)
}