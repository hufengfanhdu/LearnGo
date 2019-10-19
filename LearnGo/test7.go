//Method 结构方法集
package main

import (
	"fmt"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

//不传入 Struct pointer 只读取值
func (m Movie) Summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

//传入 Struct pointer 修改值
func (m *Movie) Modify(name string, rating float64) {
	m.Name = name
	m.Rating = rating
}

func main() {
	m := Movie{
		Name:   "Frank's Life",
		Rating: 98.30,
	}
	// fmt.Println(m.Summary())
	fmt.Println(m)
	m.Modify("Bob", 88.90)
	fmt.Println(m)
}
