package main

import "fmt"

func main() {
	fmt.Println("************************切片****************************")
	slice := make([]string, 3);
	slice[0] = "a"
	slice[1] = "b"
	slice[2] = "c"
	fmt.Println("slice:", slice)

	slice = append(slice, "d", "e", "f")
	fmt.Println("slice after appending:", slice)

	slice_copied := make([]string, len(slice))
	copy(slice_copied, slice)
	fmt.Println("slice copied from slice:", slice_copied)

	slice_s := slice[2:5]
	fmt.Println("sliced from slice with index 2 to 4:", slice_s)
	slice_s1 := slice[2:]
	fmt.Println("sliced from slice with index 2 to the last:", slice_s1)
	slice_s2 := slice[:5]
	fmt.Println("sliced from slice with index first to 4:", slice_s2)
}