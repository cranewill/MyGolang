package main

import "fmt"

func main() {
	fmt.Println("************************for****************************")
	for {
		fmt.Println("无条件for循环等同于while(true)，这里执行一次就break")
		break
	}

	for i:=0; i<3; i++ {
		fmt.Println(i)
	}
}