package main

import "fmt"

func main() {
	fmt.Println("************************pointer****************************")
	i := 10
	fmt.Println("initialize i:", i)

	valueChange(i)
	fmt.Println("after valueChange, i:", i)

	pointerChange(&i)
	fmt.Println("after pointerChange, i:", i)

	fmt.Println("i's address:", &i)
}

// 传递值
func valueChange(param int) {
	param = 0
}

// 传递指针
func pointerChange(param *int) {
	*param = 0
}
