package main

import "fmt"

// 定义一个只能用于发送消息的队列
func ping(pingChan chan<- string) {
	pingChan <- "ping"
}

func pong(pingChan <-chan string, pongChan chan<- string) {
	msg := <-pingChan
	pongChan <- msg
}

func main() {
	fmt.Println("************************channel direction****************************")

	pingChan := make(chan string, 1)
	pongChan := make(chan string, 1)

	ping(pingChan)
	pong(pingChan, pongChan)
	fmt.Println(<-pongChan)
}