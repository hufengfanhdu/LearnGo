//Struct

package main

import (
	"fmt"
	// "reflect"
)

type Ticket struct {
	Name    string
	Price   float64
	Address Address
}

type Address struct {
	Street   string
	City     string
	PostCode string
}

type Test struct {
	A string
	b string
}

//使用构造函数自定义默认值	golang并不支持函数默认值，容易产生歧义
func NewAddress(street string, city string, postcode string) Address {
	a := Address{
		Street:   street,
		City:     city,
		PostCode: "DB511",
	}
	return a
}

func main() {
	/*
	 *	结构体赋值
	 */
	// var t Ticket
	// t = Ticket{
	// 	Name:  "Hello",
	// 	Price: 200.50,
	// 	Address: Address{
	// 		Street:   "No.1 St",
	// 		City:     "HZ",
	// 		PostCode: "DB311",
	// 	},
	// }
	// fmt.Printf("%+v\n", t)

	/*
	 *	使用relect.TypeOf() 查看结构体类型
	 */
	// add1 := NewAddress("No.1", "HZ", "DB511")
	// add2 := Address{
	// 	Street:   "No.2",
	// 	City:     "BJ",
	// 	PostCode: "DB811",
	// }
	// fmt.Printf("%+v\n", add1)
	// fmt.Println(reflect.TypeOf(add1))
	// fmt.Printf("%+v\n", add2)
	// fmt.Println(reflect.TypeOf(add2))

	test := Test{
		A: "a",
		b: "b",
	}
	test2 := &test
	fmt.Printf("%+v\n", test)
	fmt.Printf("%p\n", &test)
	test2.A = "A"
	fmt.Printf("%+v\n", *test2)
	fmt.Printf("%p\n", test2)
}
