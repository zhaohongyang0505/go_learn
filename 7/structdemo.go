package main

import (
	"fmt"
	"learn/objtest"
	"learn/utils"
	"log"
)

type Liters float64
type Gallons float64
type MyType string
type Number int

func (m MyType) sayHi() {
	fmt.Println("Hi", m)
}
func (n *Number) Double() {
	*n *= 2
}
func main() {
	//struct
	var myStruct1 utils.MyStruct
	myStruct1.Number = 1.1
	myStruct1.Word = "pie"
	myStruct1.Toggle = true
	fmt.Println(myStruct1)
	//struct
	var myStruct2 utils.MyStruct
	myStruct2.Number = 1.12
	myStruct2.Word = "pie2"
	myStruct2.Toggle = false
	fmt.Println(myStruct2)
	//传递指针优点：struct很大时，不需要内存复制
	showInfo(&myStruct2)
	myStruct2.Address.City = "city1"
	myStruct2.Sheng = "sheng"
	fmt.Println(myStruct2)
	//底层数据类型的定义类型，用途：更明确的含义，单位等信息
	carFuel := Gallons(10.0)
	busFuel := Liters(240.0)
	fmt.Println(carFuel, busFuel)
	fmt.Println(Gallons(10.0) + Gallons(10.0))
	//方法与函数的区别：方法有接收器
	value1 := MyType("value1")
	value1.sayHi()
	value2 := MyType("value2")
	value2.sayHi()
	//接收器指针
	number := Number(4)
	//保存变量中，go可以自动把值转换为指针，直接用值类型不保存变量无法转换
	number.Double()
	fmt.Println(number)

	date := objtest.Date{}
	// 使用setter方法校验，防止绕过：另外的包+字段不大写+Setter大写
	// err := date.SetYear(0)
	err := date.SetYear(2024)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())

}
func showInfo(input *utils.MyStruct) {
	input.Word = "XXX"
	input.HomeAddress.City = "city"
	fmt.Println(input)
}
