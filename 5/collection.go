package main

import (
	"fmt"
	"sort"
)

// 数据结构
func main() {

	//数组
	//默认0值
	var nums [3]int
	nums[0]++
	nums[0]++
	nums[2]++
	fmt.Println(nums)
	var nums1 [3]string = [3]string{"a", "b", "c"}
	fmt.Println(nums1)
	nums2 := [3]string{"d", "e", "f"}
	fmt.Println(nums2)
	//换行
	nums3 := [3]string{
		"ddddddddddddddddddddddddddddddd",
		"eeeeeeeeeeeeeeeeeeeeeeeeeeee",
		"ffffffffffffffffffffffff",
	}
	fmt.Println(nums3)

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	for index, num := range nums3 {
		fmt.Println(index, num)
	}
	for _, num := range nums3 {
		fmt.Println(num)
	}
	for index, _ := range nums3 {
		fmt.Println(index)
	}

	//切片，不指定大小
	mySlice := make([]string, 7)
	mySlice[0] = "11"
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	mySlice = []string{"1", "2", "3", "4", "5", "6", "7"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:5])
	fmt.Println(mySlice[:3])
	fmt.Println(mySlice[1:])
	//从数组创建的切片，修改会作用于底层数组和所有切片
	testSlice := mySlice[2:5]
	fmt.Println(testSlice)
	testSlice[0] = "s"
	fmt.Println(mySlice)
	fmt.Println(testSlice)
	//append
	testSlice = append(testSlice, "d", "e")
	fmt.Println(testSlice)
	//切片变量默认值为nil，不是0值
	var intSlice []int
	fmt.Printf("intSlice:%#v\n", intSlice)
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 1)
	fmt.Println(len(intSlice))

	//可变长参数函数
	severalInts(1)
	severalInts(1, 2, 3)
	severalInts()
	several(true, 1, 2, 3)
	//像可变长参数传递切片
	severalInts(intSlice...)

	//map
	ranks := make(map[string]int)
	ranks["xxxx"] = 1
	fmt.Println(ranks)
	fmt.Println(ranks["xxxx"])

	myMap := map[string]float64{"a": 1.2, "b": 1.4}
	value, ok := myMap["a"]
	fmt.Println(value, ok)
	value, ok = myMap["c"]
	fmt.Println(value, ok)
	delete(myMap, "a")
	value, ok = myMap["a"]
	fmt.Println(value, ok)
	myMap["x"] = 2.1
	myMap["xx"] = 2.2
	myMap["xxx"] = 2.3
	for key, value := range myMap {
		fmt.Println(value, key)
	}
	//有序遍历map
	var names []string
	for key := range myMap {
		names = append(names, key)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name)
	}

}

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}
func several(flag bool, numbers ...int) {
	fmt.Println(flag, numbers)
}
