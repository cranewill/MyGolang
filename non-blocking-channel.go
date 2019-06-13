package main

import "fmt"

func main() {
	fmt.Println("************************non-blocking-channel****************************")

	messages := make(chan string)
	signals := make(chan string)

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Sent msg:", msg)
	case sig := <-signals:
		fmt.Println("Received msg:", sig)
	default:
		fmt.Println("no message")
	}
}
