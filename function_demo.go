package main

import (
	"errors"
	"fmt"
)

// 基本函数
func add(a int, b int) int {
	return a + b
}

// 简写
func add2(a, b int) int {
	return a + b
}

// 返回多值
func dlv(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("by zero")
	}
	return a / b, nil
}

// 命名返回参数
func cal(a, b int) (sum, product int) {
	sum = a + b
	product = a * b
	return
}

func sum(numbers ...int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

func testFunctionDemo() {
	fmt.Printf("%d\n", sum(1, 2))
}
