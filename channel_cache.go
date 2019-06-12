package main

import "fmt"

func main() {
	fmt.Println("************************channel cache****************************")

	// 创建一个可以缓存两个数据的通道
	cacheChannel := make(chan string, 2)

	cacheChannel <- "data_1"
	cacheChannel <- "data_2"
	
	receive1 := <- cacheChannel
	receive2 := <- cacheChannel

	fmt.Println(receive1)
	fmt.Println(receive2)
}