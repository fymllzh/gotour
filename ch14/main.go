package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//func main() {
//
//	ss:=[]string{"飞雪无情","张三"}
//
//	fmt.Println("切片ss长度为",len(ss),",容量为",cap(ss))
//
//	ss=append(ss,"李四","王五")
//
//	fmt.Println("切片ss长度为",len(ss),",容量为",cap(ss))
//
//	fmt.Println(ss)
//
//}

func main() {

	a1:=[2]string{"飞雪无情","张三"}

	s1:=a1[0:1]

	s2:=a1[:]
	fmt.Println(s1,s2)
	//打印出s1和s2的Data值，是一样的

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)

	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)
	s:="飞雪无情"

	b:=[]byte(s)

	s3:=string(b)

	fmt.Println(s,b,s3)


}
