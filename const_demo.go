package main

import "fmt"

const a int = 5

const b, c, d = 1, 2, 3

const (
	e = iota
	f
	g
)

type Gender int

const (
	Male Gender = iota
	Female
)

func (g *Gender) isMal() bool {
	return *g == Male
}

func (g *Gender) isFemale() bool {
	return *g == Female
}

func testConstDemo() {
	fmt.Println(a, b, c, d, e, f, g)
	var g Gender = Male
	fmt.Println(g.isMal())
}
