package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	default:
		panic("unhandled default case")
	}
}

func main() {
	//var a *float32 // 指针
	//var b myInt    // 自定义类型
	//var c rune     // 类型别名
	//reflectType(a) // type: kind:ptr
	//reflectType(b) // type:myInt kind:int64
	//reflectType(c) // type:int32 kind:int32
	//
	//type person struct {
	//	name string
	//	age  int
	//}
	//type book struct{ title string }
	//var d = person{
	//	name: "沙河小王子",
	//	age:  18,
	//}
	//var e = book{title: "《跟小王子学Go语言》"}
	//reflectType(d) // type:person kind:struct
	//reflectType(e) // type:book kind:struct

	//var a float32 = 3.14
	//var b int64 = 100
	//reflectValue(a) // type is float32, value is 3.140000
	//reflectValue(b) // type is int64, value is 100
	//// 将int类型的原始值转换为reflect.Value类型
	//c := reflect.ValueOf(10)
	//fmt.Printf("type c :%T\n", c) // type c :reflect.Value

	//stu1 := student{
	//	Name:  "小王子",
	//	Score: 90,
	//}
	//
	//t := reflect.TypeOf(stu1)
	//fmt.Println(t.Name(), t.Kind()) // student struct
	//// 通过for循环遍历结构体的所有字段信息
	//for i := 0; i < t.NumField(); i++ {
	//	field := t.Field(i)
	//	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	//}
	//
	//// 通过字段名获取指定结构体字段信息
	//if scoreField, ok := t.FieldByName("Score"); ok {
	//	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	//}
	s := student{}
	printMethod(s)

}

func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		fmt.Print("输出:")
		fmt.Println(args)
		v.Method(i).Call(args)
	}
}
