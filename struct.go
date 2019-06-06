package main

import "fmt"

// 声明结构体
type Person struct {
	name string
	age int
}

func main() {
	fmt.Println("************************struct****************************")
	person := Person{"will", 12}
	fmt.Println(person)
	fmt.Println("person named:", person.name)
	fmt.Println("his age is:", person.age)
	person.name = "Tsuru"
	fmt.Println("after name changing his name becomes:", person.name)
}

