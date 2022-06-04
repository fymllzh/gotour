package main

import (
	"fmt"
	"reflect"
)

//func main() {
//	var sp *string
//	sp = new(string)//关键点
//	*sp = "飞雪无情"
//	fmt.Println(sp,*sp)
//}
//func main() {
//	i:=3
//	iv:=reflect.ValueOf(i)
//	it:=reflect.TypeOf(i)
//	fmt.Println(iv,it)//3 int
//}
//func main() {
//	p:=person{Name: "飞雪无情",Age: 20}
//	pt:=reflect.TypeOf(p)
//	//遍历person的字段
//	for i:=0;i<pt.NumField();i++{
//		fmt.Println("字段：",pt.Field(i).Name)
//	}
//	//遍历person的方法
//	for i:=0;i<pt.NumMethod();i++{
//		fmt.Println("方法：",pt.Method(i).Name)
//	}
//}
//func main() {
//
//	p:=person{Name: "飞雪无情",Age: 20}
//
//	pt:=reflect.TypeOf(p)
//
//	stringerType:=reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
//
//	writerType:=reflect.TypeOf((*io.Writer)(nil)).Elem()
//
//	fmt.Println("是否实现了fmt.Stringer：",pt.Implements(stringerType))
//
//	fmt.Println("是否实现了io.Writer：",pt.Implements(writerType))
//
//}
//func main() {
//	p:=person{Name: "飞雪无情",Age: 20}
//	//struct to json
//	jsonB,err:=json.Marshal(p)
//	if err==nil {
//		fmt.Println(string(jsonB))
//	}
//	//json to struct
//	respJSON:="{\"name\":\"李四\",\"age\":40}"
//	json.Unmarshal([]byte(respJSON),&p)
//	fmt.Println(p)
//}

type person struct {

	Name string `json:"name"`

	Age int `json:"age"`

}
func (p person) String() string{
	return fmt.Sprintf("Name is %s,Age is %d",p.Name,p.Age)
}
//func main() {
//
//	p:=person{Name: "飞雪无情111",Age: 202}
//
//	pv:=reflect.ValueOf(p)
//
//	pt:=reflect.TypeOf(p)
//
//	//自己实现的struct to json
//
//	jsonBuilder:=strings.Builder{}
//
//	jsonBuilder.WriteString("{")
//
//	num:=pt.NumField()
//
//	for i:=0;i<num;i++{
//
//		jsonTag:=pt.Field(i).Tag.Get("json") //获取json tag
//
//		jsonBuilder.WriteString("\""+jsonTag+"\"")
//
//		jsonBuilder.WriteString(":")
//
//		//获取字段的值
//
//		jsonBuilder.WriteString(fmt.Sprintf("\"%v\"",pv.Field(i)))
//
//		if i<num-1{
//
//			jsonBuilder.WriteString(",")
//
//		}
//
//	}
//
//	jsonBuilder.WriteString("}")
//
//	fmt.Println(jsonBuilder.String())//打印json字符串
//
//}
func main() {
	p:=person{Name: "飞雪无情",Age: 20}
	pv:=reflect.ValueOf(p)
	//反射调用person的Print方法
	mPrint:=pv.MethodByName("Print")
	args:=[]reflect.Value{reflect.ValueOf("登录")}
	mPrint.Call(args)
}
func (p person) Print(prefix string){
	fmt.Printf("%s:Name is %s,Age is %d\n",prefix,p.Name,p.Age)
}
