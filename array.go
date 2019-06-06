package main

import "fmt"

func main() {
	fmt.Println("************************数组****************************")
	var array [5]int
	for i:=0; i < len(array); i++ {
		array[i] = i + 1
	}
	fmt.Println(array)
}