//数据类型
//bool	true or false (不能使用1 or 0)	default = false
//package relect	查看数据类型 reflect.TypeOf()
//package strconv	"true" or "false" 转 bool strconv.ParseBool

//函数
//不定参数 func functionName(numbers ... int) 使用...

package main

import (
	"fmt"
	// "reflect"
	// "strconv"
)

func functionInt(a *int) {
	fmt.Println("This is a:", a)
	fmt.Println("And this is *a:", *a)
}

func functionStr(str *string) {
	fmt.Println("This is str:", str)
	fmt.Println("And this is *str", *str)
}

//具名返回值函数
func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}

//函数作值传递
func anotherFunction(f func() interface{}) interface{} {
	return f()
}

//recurrence funciton
func functionRe(i int) int {
	if i > 10 {
		fmt.Println("All be done")
		return i
	}
	fmt.Println("Show one time")
	i += 1
	return functionRe(i)
}

func main() {

	/*
	 *  str 转 bool
	 */
	// var s string = "true"
	// b, _ := strconv.ParseBool(s)
	// fmt.Println(reflect.TypeOf(b))
	// fmt.Println(b)

	/*
	 *  str 转 int
	 */
	// var s string = "1234"
	// b, _ := strconv.Atoi(s)
	// c, _ := strconv.ParseUint(s, 10, 0)
	// fmt.Println(reflect.TypeOf(b))
	// fmt.Println(b)
	// fmt.Println(reflect.TypeOf(c))
	// fmt.Println(c)

	/*
	 *  引用以及指针测试代码
	 */
	// var a int = 1
	// var str string = "abc"
	// functionInt(&a)
	// functionStr(&str)

	/*
	 *  具名函数
	 */
	// x, y := sayHi()
	// fmt.Println("X:", x, " Y:", y)

	/*
	 *  函数作值传递
	 */
	// fn1 := func() string {
	// 	return "This is function1"
	// }
	// fn2 := func() int {
	// 	return 1
	// }
	// fmt.Println(fn1())
	// fmt.Println(fn2())

	/*
	 *  递归函数
	 */
	// functionRe(0)

}
