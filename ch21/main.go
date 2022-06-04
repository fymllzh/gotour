package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func init()  {
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetPrefix("【lizhihao】")
}

//func main() {
//	log.Println("飞雪无情的博客:","http://www.flysnow.org")
//	log.Printf("飞雪无情的微信公众号：%s\n","flysnow_org")
//	defer func() {
//	if err := recover(); err!= nil{
//			fmt.Println("bbbb")
//		}
//	}()
//	log.Fatal("hahaha")
//	fmt.Println("sodu")
//}
func main() {
	//定义零值Buffer类型变量b
	var b bytes.Buffer
	//使用Write方法为写入字符串
	b.Write([]byte("你好1oA"))
	//这个是把一个字符串拼接到Buffer里
	fmt.Fprint(&b,",","http://www.flysnow.org")
	//把Buffer里的内容打印到终端控制台
	b.WriteTo(os.Stdout)
}
