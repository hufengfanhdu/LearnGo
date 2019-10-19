package main

import (
	//S "./Struct"
	"fmt"
)

func main() {
	test := S.Test{
		A: "a",
		//b: "b",
	}
	fmt.Println(test)
	fmt.Println(test.A)
	//因为b为私有属性所以不能访问
	//fmt.Println(test.b)
}
