package main

import "fmt"

type People struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

//字符串转结构体
//func main ()  {
//	P := People{}
//	err := json.Unmarshal([]byte(`{"name":"小号","age":18}`),&P)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//	fmt.Printf("res %T",P)
//}
func NewPeople(name string, age int) *People {
	return &People{Name:name, Age:age}
}

//结构体的工厂函数
func main()  {
	aa := NewPeople("li",15)
	fmt.Printf("res %+v",aa)
}