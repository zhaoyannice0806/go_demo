package main

import "fmt"

// ============Demo 01================
type Animal interface {
	Bark() string
}

type Dog struct {
	Name string
}

func (dog Dog) Bark() string {
	return dog.Name + ": 汪汪"
}

type Cat struct {
	Name string
}

func (cat Cat) Bark() string {
	return cat.Name + ": 喵喵"
}

func makeSound(animal Animal) {
	fmt.Println(animal.Bark())
}

// ============Demo 02================
type Runner interface {
	run()
}

type Eater interface {
	eat()
}

type Pseron struct {
	Name string
}

func (p Person) run() {
	fmt.Println(p.Name + "在跑步")
}

func (p Person) eat() {
	fmt.Println(p.Name + "在吃饭")
}

// ============Demo 03================
func PrintfAnyValue(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Printf("int 类型:%d\n", val)
	case string:
		fmt.Printf("string 类型:%s\n", val)
	case bool:
		fmt.Printf("bool 类型:%t\n", val)
	default:
		fmt.Printf("未知类型:%T\n", val)
	}
}

func testInterfaceDemo() {
	dog := Dog{Name: "旺财"}
	cat := Cat{Name: "咪咪"}

	makeSound(dog)
	makeSound(cat)

	p := Person{Name: "张三"}

	var runer Runner = p
	runer.run()

	var eater Eater = p
	eater.eat()

	fmt.Printf("%p,%p,%p", &runer, &eater, &p)

	PrintfAnyValue(100)
	PrintfAnyValue("AA")
	PrintfAnyValue(true)
	PrintfAnyValue(3.1555)
}
