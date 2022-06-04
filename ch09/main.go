package main

import (
	"fmt"
	"sync"
	"time"
)

//共享的资源

var sum = 0
var mutex sync.RWMutex

func main() {

	//开启100个协程让sum+10

	for i := 0; i < 100; i++ {

		go add(10)

	}

	//防止提前退出

	time.Sleep(2 * time.Second)

	fmt.Println("和为:",sum)

}

func add(i int) {


	mutex.Lock()
	sum += i
	mutex.Unlock()

}
func readSum() int {
	//只获取读锁
	mutex.RLock()
	defer mutex.RUnlock()
	b:=sum
	return b
}
