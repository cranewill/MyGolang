package main

import "fmt"

func main() {
	fmt.Println("************************switch****************************")
	switch_i := 1
	switch switch_i {
	case 1:
		fmt.Println("it's 1")
	default :
		fmt.Println("it's not 1")
	}
}