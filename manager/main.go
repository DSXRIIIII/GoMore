package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"manager/model"
	"manager/service"
	"os"
)

func main() {
	customersJsonBytes, _ := ioutil.ReadFile("db/customers.json")
	var customers []model.Customer
	_ = json.Unmarshal(customersJsonBytes, &customers)
	var initCustomerID = 0
	if len(customers) != 0 {
		initCustomerID = customers[len(customers)-1].Cid
	}
	cs := service.CustomerService{
		Customers:  customers,
		CustomerId: initCustomerID,
	}
	for {
		fmt.Printf("\033[1;30;42m%s\033[0m\n", `
----------------客户信息管理系统--------------
		   1、添加客户
		   2、查看客户
		   3、更新客户
		   4、删除客户
           5、保存
		   6、退出
-------------------------------------------
`)
		var choice int
		fmt.Print("请输入您的选择：")
		_, _ = fmt.Scanln(&choice)
		switch choice {
		case 1:
			cs.AddCustomer()
		case 2:
			cs.CheckCustomer()
		case 3:
			cs.UpdateCustomer()
		case 4:
			cs.DeleteCustomer()
		case 5:
			fmt.Println("保存")
		case 6:
			fmt.Println("退出")
			os.Exit(0)
		}
	}
}
