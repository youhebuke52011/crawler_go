package main

import "fmt"

//type MyInterface interface {
//	Print()
//}
//
//type MyStruct struct{}
//func (ms MyStruct) Print() {}
//
//func main() {
//	x := 1
//	var y interface{} = x
//	var s MyStruct
//	var z MyInterface = s
//	fmt.Println(y, z)
//}

func Foo(x interface{}) {

	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	fmt.Println(x==nil)
	Foo(x)
}