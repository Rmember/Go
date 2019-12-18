package main

import (
	"fmt"
	"unsafe"
)
type Person struct {
Name string
Age  int
Sex  int
}

func main() {
	person := &Person{"Tim", 10 , 1}
	size := unsafe.Sizeof(person) //查看函数体占用多少内存
	fmt.Println(person.Sex,size)
}
