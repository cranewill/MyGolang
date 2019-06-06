package main

import "fmt"

func main() {
	fmt.Println("************************关联数组****************************")
	_map := make(map[string]int)
	_map["key1"] = 1
	_map["key2"] = 2 // 等同于 _map := map[string]int{"key1":1, "key2":2}
	fmt.Println("map:", _map)

	map_value1 := _map["key1"];
	fmt.Println("value1 of map:", map_value1)

	delete(_map, "key1")
	fmt.Println("map after deleting key1:", _map)

	map_value2, isKey2InMap := _map["key2"] // 可通过这种方式接收返回值大于个的方法的返回值
	fmt.Println("value2 of map:", map_value2)
	fmt.Println("is key2 in the map:", isKey2InMap)
}