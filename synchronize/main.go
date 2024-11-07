package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	x       int64
	wg      sync.WaitGroup // 等待组
	m       sync.Mutex     // 互斥锁
	lockMap sync.Map
)

func add() {
	for i := 0; i < 5000; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	wg.Done()
}

func main() {
	//wg.Add(2)
	//go add()
	//go add()
	//wg.Wait()
	//fmt.Println(x)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			lockMap.Store(key, i)
			value, _ := lockMap.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

type singleton struct{}

var instance *singleton
var once sync.Once

func getInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
