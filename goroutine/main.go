package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

//func hello() {
//	fmt.Println("hello")
//	wg.Done()
//}

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}

func main() {
	//for i := 0; i < 10; i++ {
	//	wg.Add(1) // 启动一个goroutine就登记+1
	//	go hello(i)
	//}
	//wg.Wait() // 等待所有登记的goroutine都结束

	//for i := 0; i < 5; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}

}
