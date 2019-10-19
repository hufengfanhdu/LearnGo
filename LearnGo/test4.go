//map
// 创建映射 make(map[k_type]v_type)
// 删除元素 delete(mapName, "key")
package main

import (
	"fmt"
)

func main() {
	var grades = make(map[string]int)
	grades["Math"] = 92
	grades["English"] = 95
	fmt.Println(grades)
}
