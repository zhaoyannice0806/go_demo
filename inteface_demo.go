package main

import (
	"fmt"
)

// ============Demo 01 基础语法================
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

// ============Demo 02 隐式实现================
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

// ============Demo 03 空接口================
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

// ============Demo 03 接口嵌套================

type Reader interface {
	read()
}

type Writer interface {
	write()
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	path string
}

func (f File) read() {
	fmt.Println(f.path + ": File.read")
}

func (f File) write() {
	fmt.Println(f.path + ": File.write")
}

// ============Demo 04 接口断言================

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
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

	var rw File = File{path: "xxx"}
	rw.read()
	rw.write()

	circle := Circle{radius: 5.0}

	switch shape := interface{}(circle).(type) {
	case Circle:
		fmt.Println("Circle:", shape.Area())
	case Rectangle:
		fmt.Println("Rectangle:", shape.Area())
	default:
		fmt.Println("Unknown shape")
	}

	rectangle := Rectangle{width: 5.0, height: 10.0}

	switch shape := interface{}(rectangle).(type) {
	case Circle:
		fmt.Println("Circle:", shape.Area())
	case Rectangle:
		fmt.Println("Rectangle:", shape.Area())
	default:
		fmt.Println("Unknown shape")
	}
}
