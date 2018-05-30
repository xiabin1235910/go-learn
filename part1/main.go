package main

import (
	"fmt"
	"math/rand"
	"math"
	"runtime"
)

const (
	Small = 123
	Big = 456
)

func main() {
	var ci, cool, python bool
	fmt.Println(ci, cool, python)

	vv := 123
//	vv = "sfdasf"
	fmt.Printf("%T is the type\n", vv)

	const World = "world .... "
	fmt.Println(World, Small, Big)

	lim := 5
	if lim < 6 {
		fmt.Println("This is my house ... ")
	}

	fmt.Println("My favorite number is ", rand.Intn(100))
	fmt.Printf("you have %g problems.\n", math.Sqrt(7))
	fmt.Println(add(43, 16))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d := split(17)
	fmt.Println(c, d)

	switchTest()

	pointerTest()

	structTest()

	mapTest()

	interfaceTest()

	floatTest()
}

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 2;
	y = sum - x;
	return
}

func switchTest () {
	defer fmt.Println("zzzjjj")
	
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	default: 
		fmt.Println("Default")
	}
}

func pointerTest () {
	i := 1

	fmt.Println("pointer test", i)

	p := &i

	*p = 111

	fmt.Println("pointer test for II is ", i)
}

type Vertex struct {
	X int
	Y int
}

func structTest () {
	v := Vertex{1 ,2}

	v.X = 3

	fmt.Println("Vertex is ", v)

	p := &v

	p.Y = 66

	fmt.Println("Vertex II is ", *p)

	v.Abs()
}

func (v Vertex) Abs() {
	v.X = 10086

	fmt.Println("hello ", v)
}

func mapTest() {
	m := make(map[string]int)

	m["ben"] = 1
	m["zoe"] = 2

	fmt.Println(m)
}

type Abser interface {
	Abs() 
}

func interfaceTest() {
	var a Abser

	v := Vertex{3, 4}

	a = &v

	a.Abs()
}

type F float64
func floatTest() {
	i := F(math.Pi)

	fmt.Println(i)
}
