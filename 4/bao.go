package main

import (
	"fmt"
	"learn/utils"
	"learn/utils/dates"
)

// 自定义包
func main() {
	utils.Hello()
	utils.Hi()

	//常量
	fmt.Println(dates.DaysInWeek)
}
