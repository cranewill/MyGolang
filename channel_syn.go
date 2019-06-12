package main

import "fmt"
import "time"

// 协程中执行的方法，需要一个synChan通道来向主过程同步状态
func myCoroutine(synChan chan bool) {
	fmt.Println("Start myCoroutine...")
	time.Sleep(time.Second)
	fmt.Println("done!")

	synChan <- true
}

func main() {
	fmt.Println("************************channel cache****************************")

	synChan := make(chan bool)
	go myCoroutine(synChan)

	<- synChan	// 如果没有这里取出synChan通道内消息这一句，则myCoroutine可能来不及执行主过程就已经结束了
}