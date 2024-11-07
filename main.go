package main

var number = 456

const (
	a = iota + 4
	b = iota + 5
	c = iota + 2
	d
)

func main() {
	//fmt.Println("Go Space")
	//var x int8 = 10
	//var y int16 = 20
	//fmt.Println(x + int8(y))
	//api.Add()

	//var sb string = "我是sb字符串"
	//var value int = 1
	//var value2 int = 2
	//fmt.Println(value + value2)
	//var name string
	//fmt.Scanf("%s", &name)
	//fmt.Println(name)

	//var (
	//	name = "lihh"
	//	age  = 18
	//)
	//fmt.Println(age, name)

	//number := 88
	//if true {
	//	number = 123
	//	fmt.Println(number)
	//}
	//fmt.Println(number)

	//for {
	//	fmt.Printf("%d\n", number)
	//}

	//for i := 0; i < 12; i++ {
	//	fmt.Println(i)
	//}

	//index := 12
	//index--
	//fmt.Println(index)

	//var v1 int8 = 10
	//var v2 int16 = 18
	//v3 := int16(v1) + v2
	//fmt.Println(v3, reflect.TypeOf(v3))

	//var v1 int16 = 130
	//v2 := int8(v1)
	//fmt.Println(v2)

	//v1 := 19
	//result := strconv.Itoa(v1)
	//fmt.Println(result, reflect.TypeOf(result))
	//var v2 int8 = 17
	//data := strconv.Itoa(int(v2))
	//fmt.Println(data, reflect.TypeOf(data))

	// 字符串转换为整型：转换后是int类型;可能存在错误

	//v1 := "666"
	//result, err := strconv.Atoi(v1)
	//if err == nil {
	//	fmt.Println("转换成功", result, reflect.TypeOf(result))
	//} else {
	//	fmt.Println("转换失败")
	//}

	//var name string = "李辉煌"
	//fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	//fmt.Println(len(name))

	//byteSet := []byte(name)
	//fmt.Println(string(byteSet))
	//
	//tempSet := []rune(name)
	//fmt.Println(tempSet)
	//fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	//fmt.Println(tempSet[1], strconv.FormatInt(int64(tempSet[1]), 16))
	//fmt.Println(tempSet[2], strconv.FormatInt(int64(tempSet[2]), 16))
	//
	//fmt.Println(utf8.RuneCountInString(name))
	//fmt.Println(len(name))
	//fmt.Println(strings.Contains(name, "李"))
	//
	//var sb strings.Builder
	//sb.WriteString("你是")
	//sb.WriteString("谁")
	//fmt.Println(sb.String())

	//num := "666"
	//data, _ := strconv.Atoi(num)
	//fmt.Println(data)
	//info, _ := strconv.ParseInt(num, 2, 64)
	//fmt.Println(info)

	//var name string = "李辉煌"
	//fmt.Println(name[3:6])

	//var numbers [3]int
	//numbers[0] = 1
	//numbers[1] = 2
	//numbers[2] = 3
	//fmt.Println(numbers)
	//
	//names := [...]int{1, 2, 3}
	//fmt.Println(names)

	//nums := [3]int8{11, 22, 33}
	//fmt.Printf("数组的内存地址：%p \n", &nums)
	//fmt.Printf("数组第1个元素的内存地址：%p \n", &nums[0])
	//fmt.Printf("数组第2个元素的内存地址：%p \n", &nums[1])
	//fmt.Printf("数组第3个元素的内存地址：%p \n", &nums[2])

	//nums := &[3]int{1, 2, 3}
	//(*nums)[0] = 10
	//fmt.Println(*nums)

	//nums := [3]int{1, 2, 3}
	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(i, nums[i])
	//}
	//for key, item := range nums {
	//	fmt.Println(key, item)
	//}
	//for _, item := range nums {
	//	fmt.Println(item)
	//}

	//names := [3]string{"Alex", "超级无敌小jj", "傻儿子"}
	////a. 请根据索引获取“傻儿子”
	//fmt.Println(names[2])
	////b. 请根据索引获取“alex”
	//fmt.Println(strings.ToLower(names[0]))
	////c. 请根据索引获取“小jj”
	//fmt.Println(names[1][12:17])
	////d. 请将name数组的最后一个元素修改为 “大烧饼”
	//names[2] = "大烧饼"
	//fmt.Println(names)

	//var nestData [3][2]int
	//fmt.Println(nestData)
	//
	//dataList := [][]string{
	//	{"alex", "qwe123"},
	//	{"eric", "admin11"},
	//	{"tony", "pp1111"},
	//}
	//for _, data := range dataList {
	//	fmt.Printf("我是%s,我的密码是%s。\n", data[0], data[1])
	//}

	//userList := [][]string{
	//	{"alex", "qwe123"},
	//	{"eric", "admin11"},
	//	{"tony", "pp1111"},
	//}
	//
	//var username, password string
	//fmt.Print("请输入用户名：")
	//_, _ = fmt.Scan(&username)
	//fmt.Print("请输入密码：")
	//_, _ = fmt.Scan(&password)
	//valid := false
	//for _, user := range userList {
	//	if user[0] == username && user[1] == password {
	//		valid = true
	//		break
	//	}
	//}
	//if valid {
	//	fmt.Println("用户名和密码正确。")
	//} else {
	//	fmt.Println("用户名或密码错误。")
	//}

	//var users = make([]int, 1, 3)
	//fmt.Println(len(users), cap(users))

	//v1 := make([]int, 1, 3)
	//v2 := append(v1, 66)
	//fmt.Println(v1)
	//fmt.Println(v2)
	//
	//v1[0] = 99
	//fmt.Println(v1)
	//fmt.Println(v2)
	//
	//fmt.Println(reflect.TypeOf(v1))

	//v1 := []int{11, 22, 33}
	//v2 := append(v1, 44)
	//v3 := append(v1, 55, 66, 77, 88)
	//v4 := append(v1, []int{100, 200, 300}...)
	//fmt.Println(v2)
	//fmt.Println(v3)
	//fmt.Println(v4)

	//v1 := []int{11, 22, 33, 44, 55, 66}
	//insertIndex := 3
	//result := make([]int, 0, len(v1)+1)
	//result = append(result, v1[:insertIndex]...)

	//v1 := []int{11, 22, 33, 44, 55, 66}
	//for i := 0; i < len(v1); i++ {
	//	fmt.Println(i, v1[i])
	//}
	//
	//name := "李辉煌"
	//fmt.Println(utf8.RuneCountInString(name))

	//userInfo := map[string]string{"1": "2", "3": "4"}
	//fmt.Println(userInfo["1"])

	//data := make(map[string]int)
	//data["100"] = 998
	//data["200"] = 999

	//v1 := make(map[[2]int]float32)
	//v1[[2]int{1, 1}] = 1.6
	//v1[[2]int{1, 2}] = 3.4
	//fmt.Println(v1)

	//info := make(map[string]string, 10)
	//info["n1"] = "李辉煌"
	//info["n2"] = "user"
	//
	//v7 := make(map[string][2]map[string]string)
	//v7["n1"] = [2]map[string]string{{"name": "武沛齐", "age": "18"}, {"name": "alex", "age": "78"}}
	//v7["n2"] = [2]map[string]string{{"name": "eric", "age": "18"}, {"name": "seven", "age": "78"}}

	//name := "李辉煌"
	//var pointer *string = &name
	//fmt.Println(*pointer)

	//ChangeName(&name)
	//fmt.Println(name)
	//name := "李辉煌"
	//var p1 *string = &name
	//*p1 = "张三"
	//var p2 **string = &p1
	//**p2 = "李四"
	//var p3 ***string = &p2
	//***p3 = "?"
	//fmt.Println(name)

	//dataList := [3]int8{11, 22, 33}
	//fmt.Printf("数组的地址：%p；数组第一个元素的地址：%p \n", &dataList, &dataList[0])

	//dataList := [3]int8{11, 22, 33}
	//var p1 *int8 = &dataList[0]
	//p2 := unsafe.Pointer(p1)
	////转换成uintptr类型，然后进行内存地址的计算（即：地址加1个字节，意味着取第2个索引位置的值）。
	//targetAddress := uintptr(p2) + 1
	//newPtr := unsafe.Pointer(targetAddress)
	//value := (*int8)(newPtr)
	//fmt.Println("最终结果为：", *value)

	//var person Person
	//person.city = "123"
	//fmt.Println(person)

	//type Person struct {
	//	name   string
	//	age    int
	//	hobby  [2]string
	//	num    []int
	//	parent map[string]string
	//}
	//
	//p1 := Person{
	//	name:   "二狗子",
	//	age:    19,
	//	hobby:  [2]string{"裸奔", "大保健"},                                  // 拷贝
	//	num:    []int{69, 19, 99, 38},                                   // 未拷贝 (内部维护指针指向数据存储的地方)
	//	parent: map[string]string{"father": "Alex", "mother": "Monika"}, // 未拷贝 (内部维护指针指向数据存储的地方)
	//}
	//var p2 *Person = &p1
	//p2.name = "李辉煌"
	//fmt.Println(p1.name)
	//fmt.Println(reflect.TypeOf(p1))

	//type Person struct {
	//	name string "姓名"
	//	age  int32  "年龄"
	//	blog string "博客"
	//}
	//p1 := Person{name: "李辉煌", age: 18, blog: "https://www.pythonav.com"}
	//p1Type := reflect.TypeOf(p1)
	//filed1 := p1Type.Field(0)
	//fmt.Println(filed1.Tag)
	//
	//filed2, _ := p1Type.FieldByName("blog")
	//fmt.Println(filed2.Tag)
	//filedNum := p1Type.NumField()
	//for index := 0; index < filedNum; index++ {
	//	filed := p1Type.Field(index)
	//	fmt.Println(filed.Name, filed.Tag)
	//}
	//var schoolList []School
	//for {
	//	var band, city string
	//	fmt.Printf("请输入品牌：")
	//	_, _ = fmt.Scanf("%s", &band)
	//	if band == "Q" {
	//		break
	//	}
	//	fmt.Printf("请输入城市：")
	//	_, _ = fmt.Scanf("%s", &city)
	//	// 创建一个结构体对象
	//	sch := School{band: band, city: city}
	//	// 把结构体对象放到切片中
	//	schoolList = append(schoolList, sch)
	//}
	//fmt.Println(schoolList)
	//value, b := add(1, 2)
	//fmt.Println(value, b)

	//var functionList []func()
	//
	//for i := 0; i < 5; i++ {
	//	function := func() {
	//		fmt.Println(i)
	//	}
	//	functionList = append(functionList, function)
	//}
	//
	//functionList[0]()
	//functionList[1]()
	//functionList[2]()

	//var functionList []func()
	//for i := 0; i < 5; i++ {
	//	function := func(arg int) func() {
	//		return func() {
	//			fmt.Println(arg)
	//		}
	//	}(i)
	//	functionList = append(functionList, function)
	//}
	//
	//// 运行函数
	//functionList[0]()
	//functionList[1]()
	//functionList[2]()

}

//func exec(num1 int, num2 int) string {
//	fmt.Println("执行函数了")
//	return "成功"
//}
//
//type f1 func(int, int) string
//
//func getFunction() f1 {
//	return exec
//}

//func add(num1 int, num2 int) (int, bool) {
//	result := num2 + num1
//	return result, true
//}

//type School struct {
//	band string
//	city string
//}

//type Address struct {
//	city, state string
//}
//type Person struct {
//	name    string
//	age     int
//	Address // 匿名字段，那么默认Person就包含了Address的所有字段
//}

//func ChangeName(data *string) {
//	*data = "123"
//}
