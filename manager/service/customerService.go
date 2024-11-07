package service

import (
	"fmt"
	"manager/model"
)

type CustomerService struct {
	Customers  []model.Customer
	CustomerId int
}

func (cs *CustomerService) AddCustomer() {
	fmt.Printf("\033[1;35;40m%s\033[0m\n", "---------------------------添加客户开始-----------------------------")
	var name string
	fmt.Print("请输入客户姓名：")
	_, _ = fmt.Scan(&name)

	var gender string
	fmt.Print("请输入客户性别：")
	_, _ = fmt.Scan(&gender)

	var age int
	fmt.Print("请输入客户年龄：")
	_, _ = fmt.Scan(&age)

	var email string
	fmt.Print("请输入客户邮箱：")
	_, _ = fmt.Scan(&email)
	cs.CustomerId++
	var newCustomer = model.Customer{
		Cid:    cs.CustomerId,
		Name:   name,
		Age:    age,
		Gender: gender,
		Email:  email,
	}
	cs.Customers = append(cs.Customers, newCustomer)
	fmt.Printf("\033[1;35;40m%s\033[0m\n", "---------------------------添加客户成功-----------------------------")
}

func (cs CustomerService) CheckCustomer() {
	for _, customer := range cs.Customers {
		fmt.Printf("客户编号：%s 姓名：%-6s 性别：%-6s 年龄：%-6s 邮箱：%-6s\n",
			customer.Cid,
			customer.Name,
			customer.Gender,
			customer.Age,
			customer.Email)
	}
	fmt.Printf("\033[1;32;40m%s\033[0m\n", "----------------------------------客户列表结束----------------------------")
}

func (cs CustomerService) UpdateCustomer() {
	for {
		var cur int
		fmt.Println("请输入修改的客户编号")
		_, _ = fmt.Scan(&cur)
		var updateIndex = -1
		for index, customer := range cs.Customers {
			if customer.Cid == cur {
				updateIndex = index
			}
		}

		if updateIndex == -1 {
			fmt.Println("该编号不存在")
			continue
		}
		var updateCustomer *model.Customer
		updateCustomer = &cs.Customers[updateIndex]
		var name string
		fmt.Printf("请输入客户姓名(%s)：", updateCustomer.Name)
		_, _ = fmt.Scanln(&name)

		var gender string
		fmt.Printf("请输入客户性别(%s)：", updateCustomer.Gender)
		_, _ = fmt.Scanln(&gender)

		var age int
		fmt.Printf("请输入客户年龄(%s)：", updateCustomer.Age)
		_, _ = fmt.Scanln(&age)

		var email string
		fmt.Printf("请输入客户邮箱(%s)：", updateCustomer.Email)
		_, _ = fmt.Scanln(&email)

		if name != "" {
			updateCustomer.Name = name
		}
		if gender != "" {
			updateCustomer.Gender = gender
		}
		if age != 0 {
			updateCustomer.Age = age
		}
		if email != "" {
			updateCustomer.Email = email
		}
		break
	}
}

func (cs *CustomerService) DeleteCustomer() {
	for {
		var cur int
		fmt.Println("请输入删除的客户编号")
		_, _ = fmt.Scan(&cur)
		var deleteIndex = -1
		for index, customer := range cs.Customers {
			if customer.Cid == cur {
				deleteIndex = index
			}
		}

		if deleteIndex == -1 {
			fmt.Println("该编号不存在")
			continue
		}
		customers := append(cs.Customers[:deleteIndex], cs.Customers[deleteIndex+1:]...)
		fmt.Println(customers)
		fmt.Printf("\033[1;31;40m%s\033[0m\n", "---------------------------删除客户完成----------------------")
		break
	}
}
