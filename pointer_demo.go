package main

import "fmt"

func testPointerDemo() {
	fmt.Printf("===========指针基础=============\n")
	var a int = 5
	fmt.Printf("变量a:%d,内存地址:%p\n", a, &a)

	// 声明指针
	var p *int = &a
	fmt.Printf("指针p指向地址:%p,对应地址的值:%d,p的地址:%p\n", p, *p, &p)

	// 通过指针修改指向变量值
	*p = 20
	fmt.Printf("修改后a的值:%d\n", a)

	var p1 *int
	fmt.Printf("空指针:%v\n", p1)

	fmt.Printf("===========指针值传递=============\n")

	modifyVaue(a)
	fmt.Printf("值引用修改a的值结果:%d\n", a)

	modifyVaueByPoint(&a)
	fmt.Printf("指针修改a的值结果:%d\n", a)
}

func modifyVaue(x int) {
	x = 100
}

func modifyVaueByPoint(p *int) {
	*p = 100
}
