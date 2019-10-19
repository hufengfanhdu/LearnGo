//Goroutine	携程
//缓冲通道 make(chan type, length)	如果length < 消息数量会发生死锁
//	all goroutines are asleep - deadlock!
package main

import (
	"fmt"
)

func slowFunc(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	c := make(chan string, 2)
	c <- "This is the first message"
	c <- "This is the second message"
	c <- "Test"
	close(c)
	slowFunc(c)
}
