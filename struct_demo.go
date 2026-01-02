package main

import "fmt"

// 声明结构体
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	ID string
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func (p *Person) addAge() {
	p.Age++
}

/**
结构体类型
*/
func testStructDemo() {
	// 初始化结构体
	p1 := Person{
		Name: "zhangsan",
		Age:  12,
	}

	// 简短初始化
	p2 := Person{"lisi", 15}

	// 部分初始化
	p3 := Person{Name: "wangwu"}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println(p2.GetInfo())

	p3.addAge()
	fmt.Println(p3.GetInfo())

	e1 := Employee{
		Person: Person{Name: "lisi", Age: 15},
		ID:     "E001",
	}
	fmt.Println(e1.Person)
	fmt.Println(e1.ID)
	fmt.Println(e1.Name)
}
