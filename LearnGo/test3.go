//Slice
// 创建切片 make([]type, size)
// 添加元素 append -> append为不定参数函数	slice := append(slice, new items)
// 删除元素 删除第i个item -> append(slice[:i], slice[i+1:]...)

/*
	容量(cap)和长度(len)的区别：

	在与当你用 append 扩展长度时，如果新的长度小于容量，不会更换底层数组
	否则，go 会新申请一个底层数组，拷贝这边的值过去，把原来的数组丢掉
	也就是说，容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡
*/
package main

func main() {
	/*
	 *	Slice切片的make、cap 与 append 的关系
	 *	使用append 执行 添加 与 删除 操作
	 */
	//numbers := make([]int, 1, 3)
	//numbers[0] = 1
	//fmt.Println("Before append:", numbers, "Size:", len(numbers), "Cap:", cap(numbers))
	//fmt.Printf("%p\n", &numbers)
	//numbers = append(numbers, 2)
	//fmt.Println("Before append:", numbers, "Size:", len(numbers), "Cap:", cap(numbers))
	//fmt.Printf("%p\n", &numbers)
	////当 append 超过初始化的 cap 大小时，会重新申请一个数组，cap大小为len+1
	//numbers = append(numbers, 3, 4, 5)
	//fmt.Println("Before append:", numbers, "Size:", len(numbers), "Cap:", cap(numbers))
	//fmt.Printf("%p\n", &numbers)
	////删除index为2的元素
	////下方append必须使用...否则会 报错：类型不匹配， 使用...确保number[2+1:]数组全部赋值
	//numbers = append(numbers[:2], numbers[2+1:]...)
	//fmt.Println("Before append:", numbers, "Size:", len(numbers), "Cap:", cap(numbers))
	//fmt.Printf("%p\n", &numbers)

	/*
	 *	Slice切片 copy操作
	 *	拷贝副本时会重新分配一个数组
	 */
	// str := make([]string, 2)
	// str[0] = "A"
	// str[1] = "B"
	// fmt.Println(str)
	// fmt.Printf("%p\n", str)
	// var str2 = make([]string, 2)
	// fmt.Printf("%s:%p\n", "Copy前地址", str2)
	// copy(str2, str[1:])
	// fmt.Println(str2)
	// fmt.Printf("%p\n", str2)
	// str[1] = "MODIFY"
	// fmt.Println("str:", str)
	// fmt.Println("str2:", str2)

	//test
	//var numbers = make([]int, 4, 4)
	//numbers[0] = 1
	//numbers[1] = 2
	//numbers[2] = 3
	//numbers[3] = 4
	//var numbers2 = make([]int, 2)
	//numbers2 = numbers[2:4]
	//fmt.Println(numbers)
	//fmt.Println(numbers2)
}
