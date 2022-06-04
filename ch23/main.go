package main

import (
	"fmt"
	"reflect"
)

//
//func main() {
//	u:= User{"张三",20}
//	t:=reflect.TypeOf(u)
//	v:=reflect.ValueOf(u)
//	fmt.Println(t,v)
//	u1:=v.Interface().(User)
//	fmt.Println(u1)
//	t1:=v.Type()
//	fmt.Println(t1)
//	fmt.Println(t.Kind())
//
//	for i:=0;i<t.NumField();i++ {
//		fmt.Println(t.Field(i).Name)
//	}
//	for i:=0;i<t.NumMethod() ;i++  {
//		fmt.Println(t.Method(i).Name)
//	}
//
//
//	x:=2
//	v3:=reflect.ValueOf(&x)
//	v3.Elem().SetInt(44)
//	fmt.Println(v3,x)
//}
//type User struct{
//	Name string `json namess`
//	Age int `json age123`
//}
func main() {
	u:=User{"张三",20}
	v:=reflect.ValueOf(u)

	mPrint:=v.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))

}
type User struct{
	Name string
	Age int
}
func (u User) Print(prfix string){
	fmt.Printf("%s:Name is %s,Age is %d",prfix,u.Name,u.Age)
}

