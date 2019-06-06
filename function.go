package main

import "fmt"

func main() {
	fmt.Println("************************函数****************************")
	myFuncResult := myFunc(1, 2)
	fmt.Println("myFunc result:", myFuncResult)
	
	myMultiReturnFunc_1, myMultiReturnFunc_2 := myMultiReturnFunc()
	fmt.Print("first of myMultiReturnFunc result:", myMultiReturnFunc_1)
	fmt.Println(", second of myMultiReturnFunc result:", myMultiReturnFunc_2)

	sumServeralNum(1, 2, 3, 4)

	selfAddInteger := intAdd()
	fmt.Println(selfAddInteger)
	fmt.Println(selfAddInteger)
	fmt.Println(selfAddInteger)
}

func myFunc (a int, b int) int {
	return a + b
}

// 多返回值函数
func myMultiReturnFunc () (int, int) {
	return 100, 200
}

// 变参函数
func sumServeralNum (nums... int) {
	fmt.Print("sum of ", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(" is:", sum)
}

// 闭包函数
func intAdd() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}