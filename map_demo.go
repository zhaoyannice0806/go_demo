package main

import "fmt"

/**
map 数据类型
*/
func testMapDemo() {
	var mp map[string]int     // 声明map
	mp = make(map[string]int) // 初始化

	mp2 := map[string]int{ //初始化
		"apple":  2,
		"banana": 5,
	}

	mp["apple"] = 0 // 添加元素

	// 取值
	fmt.Println(mp2["banana"])
	fmt.Println(mp["apple"])

	value := mp["apple"]
	fmt.Println(value)

	value2, ok := mp["apple2"]
	fmt.Println(value2, ok)

	mp["apple"] = 10
	fmt.Println(mp["apple"])

	// 遍历
	for k, v := range mp2 {
		fmt.Printf("%s:%d\n", k, v)
	}
}
