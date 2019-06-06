package main

import "fmt"

func main() {
	fmt.Println("************************Range遍历****************************")

	var array [5]int
	for i:=0; i < len(array); i++ {
		array[i] = i + 1
	}

	_map := make(map[string]int)
	_map["key1"] = 1
	_map["key2"] = 2

	sum := 0
	for index, element := range array { // for这里只跟一个对象，则默认是index，而不是元素，如果只想要元素而不使用index，可以for _, element
		fmt.Println(index)
		fmt.Println(element)
		sum += element
	}
	fmt.Println("sum of array elements:", sum)

	for key, value := range _map {
		fmt.Print("range of map:key:" + key + ", ")
		fmt.Println("value:", value)
	}
}