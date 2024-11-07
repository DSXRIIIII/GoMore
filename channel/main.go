package main

import "fmt"

var ch1 chan int   // 声明一个传递整型的通道
var ch2 chan bool  // 声明一个传递布尔型的通道
var ch3 chan []int // 声明一个传递int切片的通道
var ch chan int

func main() {
	//ch4 := make(chan int)
	//ch5 := make(chan bool, 1) // 声明一个缓冲区大小为1的通道
	//ch := make(chan int)
	//ch <- 10
	//x := <-ch
	//fmt.Println(x)

	//简单来说就是无缓冲的通道必须有至少一个接收方才能发送成功
	//ch := make(chan int, 1)
	//go recv(ch) // 创建一个 goroutine 从通道接收值
	//ch <- 10
	//fmt.Println("发送成功")

	//ch := make(chan int, 2)
	//ch <- 1
	////ch <- 2
	//close(ch)
	//f2(ch)

	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25
}

func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func Producer() chan int {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()
	return ch
}

func Consumer(ch chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}
