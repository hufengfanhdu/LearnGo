/*
 *	Method: go方法调用中
 *  实参为值 形参为指针时 p.Method() 会被解释为 (&p).Method()
 *	实参为指针 形参为值时 p.Method() 会被解释为 (*p).Method()
 *
 *	若使用函数
 *  实参为值 形参为指针时 Function(p) 将报错
 *	实参为指针 形参为值时 Function(p) 将报错
 */
package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	X, Y float64
}

//方法
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//函数
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v为值，非指针
	v := Vertex{3, 4}
	fmt.Println(reflect.TypeOf(v))
	fmt.Printf("%+V\n", v)
	fmt.Printf("%p\n", v)
	fmt.Println("v:", v)

	// p为指针
	p := &v
	fmt.Println(reflect.TypeOf(p))
	fmt.Printf("%+V\n", p)
	fmt.Printf("%p\n", p)
	fmt.Println("p:", *p)

	p.Scale(2)
	fmt.Println("v:", v)
	fmt.Println("p:", *p)

}
