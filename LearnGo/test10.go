//通道阻塞
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	//每秒钟生成一个 Struct
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C
	}
}


func main() {
	messages := make(chan string)
	go pinger(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
