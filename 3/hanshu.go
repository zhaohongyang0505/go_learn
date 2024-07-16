package main

import "fmt"

// 函数
func paint(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("0 error")
	}
	area := width * height
	fmt.Printf("area: %0.3f\n", area/10)
	return area / 10, nil
}

func createPoint() *float64 {
	var point = 88.8
	return &point
}
func main() {
	//格式化动词
	fmt.Printf("A float: %0.3f\n", 3.1415)
	fmt.Printf("A float: %7.3f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")

	//函数
	var amount, total float64

	amount, err := paint(4.2, -3.0)
	fmt.Println(err)
	total += amount
	amount, err = paint(5.2, 3.5)
	fmt.Println(err)
	total += amount
	fmt.Println(total)

	//指针
	//定义int变量
	var myInt int
	//定义int指针变量
	var myIntPoint *int
	//给变量分配指针
	myIntPoint = &myInt
	//打印指针
	fmt.Println(myIntPoint)
	//打印指针的值
	fmt.Println(*myIntPoint)
	//改变指针的值
	*myIntPoint = 8
	fmt.Println(myIntPoint)
	fmt.Println(*myIntPoint)

	//函数指针
	point := createPoint()
	fmt.Println(*point)

}
