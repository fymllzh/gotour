//
//package main
//import (
//"fmt"
//"errors"
//)
//
//
////此函数读取配置文件信息
////如果文件名不正确，返回自定义错误
//func readConfFile(FileName string) (err error) { //返回error类型
//	if FileName == "config.ini" {
//		return nil //表示没有错误
//	} else {
//		return errors.New("读取文件错误")
//	}
//}
//
//func error_func(){
//	err := readConfFile("config.inis")  //这里故意写错，报错代码在第一块，如果写对，在第二块
//	if err != nil {
//		//如果读取文件发生错误，就输出这个错误，并终止程序
//		panic(err)  //这个函数作用是打印错误信息，并终止程序
//	}
//	fmt.Println("error_func()继续执行")
//}
//func main() {
//	error_func()
//	fmt.Printf("发生错误后面的代码")
//}





package main
import (
"fmt"
"errors"
)


func foo()(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("再然后④\n")
			n++
			fmt.Println("最后⑤\n")
		}
	}()
	fmt.Println("首先①\n")
	n++
	fmt.Printf("然后②n=%v\n" ,n)
	fmt.Println("然后③\n")
	panic(errors.New("i'm a bug"))
	fmt.Println("没走\n")

	return n
}
func main() {
	n := foo()
	fmt.Printf("n最后的值%v", n)
}


//这里举例，在数学计算中0是不可以作为被除数的
//func error_func() {
//	//这里使用defer + recover来捕获处理异常
//	defer func() {  //defer就是把匿名函数压入到defer栈中，等到执行完毕后或者发生异常后调用匿名函数
//		err := recover()  //recover是内置函数，可以捕获到异常
//		if err != nil {   //说明有错误
//			fmt.Println("err=", err)
//			//当然这里可以把错误的详细位置发送给开发人员
//			//send email to admin
//		}
//	}()
//	num1 := 10
//	num2 := 0
//	res := num1 / num2
//	fmt.Println("res=", res)
//}
//
//func main() {
//	//这样程序不会轻易挂掉
//	error_func()
//	i := 0
//	for {
//		i++
//		fmt.Println("发生错误后面的代码", i)
//		time.Sleep(time.Second)
//	}
//
//}