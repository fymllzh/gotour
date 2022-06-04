package main

import (
	"fmt"
	"strings"
)

const (
	one = iota
	two
	three
)

func main()  {
	//var(
	//i = 10
	//bf = true
	//s1 = "haha"
	//)
	//i :="hahaha"
	//i= "kkkkkkk"
    //pi:=&i
/*	i := 1.00
	bf := float64(i)*/
	//var a int32  = 10
	//var b int64 = int64(a)
	//var c float32 = 12.3
	//var d float64 =float64(c)
	//fmt.Println(a,b,c,d)
	// 输出各数值范围
	//将整型转浮点型111
	ss := []string{"Monday", "Tuesday", "Wednesday"}
	s := strings.Join(ss, "|")
	fmt.Println(s)

}