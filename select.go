package main

import "fmt"
import "time"

func main() {
	fmt.Println("************************select****************************")

	chan_1 := make(chan string)
	chan_2 := make(chan string)
	chan_3 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		chan_1 <- "msg_1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		chan_2 <- "msg_2"
	}()

	for i := 1; i <= 2; i++ {
		// 这里使用通道选择器对不同通道收到消息进行处理
		select {
		case msg1 := <-chan_1:
			fmt.Println("Receive msg1:", msg1)
		case msg2 := <-chan_2:
			fmt.Println("Receive msg2:", msg2)
		}
	}

	// 超时处理 select默认处理第一个准备好的case，随意当超过2s后会马上处理超时事件
	go func() {
		time.Sleep(time.Second * 3)
		chan_3 <- "msg_3"
	}()

	select {
	case msg3 := <-chan_3:
		fmt.Println("Receive msg3:", msg3)
	case <-time.After(time.Second * 2):
		fmt.Println("Receive msg3 time out!")
	}
}