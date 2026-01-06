package main

import "fmt"

// defer 在ruturn之后执行
func exmple() int {
	defer fmt.Println(" world")
	fmt.Print("Hello,")
	return 33
}

// defer 在painc之后执行
func panicExmple() {
	defer fmt.Printf("defer 执行")
	fmt.Println("开始执行")
	panic("发生错误")
	fmt.Println("不会被执行")
}

// 多个defer 执行后进先出LIFO
func multipleDefer() {
	fmt.Println("函数执行开始")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("函数执行结束")
}

// defer 语句中的参数值会在 defer 语句执行时立即计算并捕获，而不是在 defer 真正执行时计算
func deferValue() {
	i := 1
	defer fmt.Println("defer 1", i)
	i++
	defer fmt.Println("defer 2", i)
	i++
	fmt.Println("函数内值", i)
}

func testDeferDemo() {
	exmple()
	// panicExmple()
	multipleDefer()

	deferValue()
}
