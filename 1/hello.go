package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello, Go!")
	//类型
	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(5.2))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(1.0))
	fmt.Println(reflect.TypeOf("hello"))

	//变量
	var num int
	fmt.Println(num)
	num = 888
	fmt.Println(num)
	var i, j float64
	i, j = 0.1, 0.4
	fmt.Println(i, j)
	//直接赋值，省略类型
	var name = "haha"
	fmt.Println(name)

	//零值
	var myInt int
	var myFloat float64
	fmt.Println(myInt, myFloat)

	var myString string
	var myBool bool
	fmt.Println(myString, myBool)

	//短变量
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"
	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	//只有名称是以大写字母开头的变量、函数或类型才被认为是可导出的：可以从当前包之外的包访问。

	//类型转换
	var price int = 100
	fmt.Println("Price is", price, "dollars.")
	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")
	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")
	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", total <= float64(availableFunds))

	//时间
	var now time.Time = time.Now()
	var year int = time.Now().Year()
	fmt.Println(now, year)

	//替换
	s := "#1#2"
	replace := strings.NewReplacer("#", "a")
	s1 := replace.Replace(s)
	fmt.Println(s1)

	//多个返回值(常见于处理异常)
	reader := bufio.NewReader(os.Stdin)
	//空白标识符
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	//至少有一个变量是新的声明，旧的为赋值
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println(status)
}
