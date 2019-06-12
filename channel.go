package main

import "fmt"

func main() {
	fmt.Println("************************channel****************************")

	// 通过make(chan [type])来创建一个通道
	myChannel := make(chan string)

	go func(msg string) {
		// 通过<-来将要传递的值交给通道
		myChannel <- msg
	}("test")

	// 通过<-从通道中取出传递的值
	receive := <- myChannel

	// 单个通道的发送和接收是阻塞的，必须双方都准备好才会发送，所以不需要额外的同步操作就可以成功接收到消息
	fmt.Println(receive)
}