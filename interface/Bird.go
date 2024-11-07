package main

import "fmt"

// Mover 定义一个接口类型
type Mover interface {
	Move()
}
type Sayer interface {
	Say()
}

// Dog 狗结构体类型
type Dog struct{}

// Move 使用值接收者定义Move方法实现Mover接口
func (d *Dog) Move() {
	fmt.Println("狗会动")
}
func (d *Dog) Say() {
	fmt.Println("狗会叫")
}

type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func main() {
	var x Mover
	//var d1 = Dog{}
	//x = d1
	//x.Move()

	var d2 = &Dog{} // d2是Dog指针类型
	d2.Say()

	var c1 = &Cat{} // c1是*Cat类型
	x = c1          // 可以将c1当成Mover类型
	x.Move()

	// 下面的代码无法通过编译
	//var c2 = Cat{} // c2是Cat类型
	//x = c2         // 不能将c2当成Mover类型
}
