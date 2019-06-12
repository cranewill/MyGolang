package main

import "fmt"

func cor_1(par string) {
	for i := 0; i < 3; i++ {
		fmt.Println(par, ":", i)
	}
}

func main() {
	fmt.Println("************************coroutine****************************")

	go cor_1("first")

	go func(par string) {
		fmt.Println(par)
	}("second")

	var input string
	fmt.Scanln(&input) // 这里用scan是为了让主过程等待上面两个协程执行完毕后再退出
	fmt.Println("done")

}