package main

import "fmt"

func main()  {
	//ch :=make(chan string)
	//go func() {
	//	time.Sleep(time.Second)
	//	fmt.Println("lizhihao")
	//	ch <- "我是goroutine"
	//}()
	//fmt.Println("wa shi main")
	//time.Sleep(time.Second)
	//v :=<-ch
	//fmt.Println("接受到管道中的字符：",v,"haha")

	cacheCh:=make(chan int,5)

	cacheCh <- 2

	cacheCh <- 3

	fmt.Println("cacheCh容量为:",cap(cacheCh),",元素个数为：",len(cacheCh))
	close(cacheCh)

}
