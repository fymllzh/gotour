package main

import "fmt"

func main() {
	array := []int{1, 2, 4, 7}
	aa := array[1:3]
	aa[0] = 9
	for i, v := range array {
		fmt.Printf("数组的索引为%d,数组的值为%d \n\r", i, v)
	}
	for i, v := range aa {
		fmt.Printf("数组的索引为%d,数组的值为%d \n\r", i, v)
	}
	slice2 := append(array, 100)
	fmt.Println(array, slice2)
	//map1 := make(map[string]int)
	map1 := map[string]int{}
	map1["li"] = 10
	fmt.Println(map1, len(map1))
	name, o := map1["lz"]
	if o {
		fmt.Println(name)
	} else {
		fmt.Println("不存在")
		s := "hello ，看i这集好"
		bs := []byte(s)
		for i, v := range bs {
			fmt.Printf("数组的索引为%d,数组的值为%d \n\r", i, v)

		}

	}
}
