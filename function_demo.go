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

// 可变参数
func sum(numbers ...int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}

// 闭包
func makeCounter() func() int {
	num := 0
	return func() int {
		num++
		return num
	}
}

// 函数柯里化
func addCurrying(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func testFunctionDemo() {
	fmt.Printf("%d\n", sum(1, 2, 4, 5, 6))

	fmt.Println("===========闭包==================")
	counter := makeCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println("===========函数柯里化==================")
	add := addCurrying(1)
	fmt.Println(add(2))
	fmt.Println(add(3))
	fmt.Println(addCurrying(4)(4))
}
