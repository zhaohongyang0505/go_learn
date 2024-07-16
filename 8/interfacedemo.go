package main

import "fmt"

type Whistle string

func (w *Whistle) MakeSound() {
	if *w == "on" {
		fmt.Println("Tweet!")
	}
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}
func (h Horn) duanyan() {
	fmt.Println(h)
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	// 接口无需明确声明实现，一个类型包含接口声明中所有的方法就可以使用接口
	s := Whistle("on")
	var toy NoiseMaker = &s
	toy.MakeSound()
	play(Horn("Toyco Blaster"))

	//类型断言
	toy = Horn("hahhaa")
	var dy Horn = toy.(Horn)
	dy.duanyan()
	//不能确认接口原类型是什么时，使用可选ok值（返回的第二个bool值）处理类型不同时的运行时异常
	//error接口，stringer接口
}
