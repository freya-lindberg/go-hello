package main

import "fmt"

func main() {
	var i = 15.5
	var txt = "Hello world"

	fmt.Printf("%v \n", i)
	fmt.Printf("%#v \n", i)
	fmt.Printf("%v%% \n", i)
	fmt.Printf("%T \n\n", i)

	fmt.Printf("%v \n", txt)
	fmt.Printf("%#v \n", txt)
	fmt.Printf("%T \n\n", txt)

	const PI float64 = 3.141592
	fmt.Printf("PI: %v %.2f\n", PI, PI, PI)
}

func t7() {
	var x int = 1
	fmt.Printf("log: %v %T\n", x, x)
}

func t6() {
	fmt.Print(true, false, "\n")
	fmt.Println(true, false)
	fmt.Print(1, 2, "\n")
	fmt.Println(1, 2)
	fmt.Print("Cherry", "Fruit", "\n")
	fmt.Println("Cherry", "Fruit")
}

func t5() {
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	a, b, c := 4, 5, 6
	fmt.Println(a, b, c)

	var xx, yy = false, "Apple"
	fmt.Println(xx, yy)

	zz, ww := 3.0, 'c'
	fmt.Println(zz, ww)

	var (
		aa int    = 1
		bb bool   = false
		cc string = "Banana"
	)
	fmt.Println(aa, bb, cc)

	const PI = 3.14
	const ONE int = 1
	fmt.Println(PI, ONE)
}

func t4() {
	var s string
	s = "Freya Lindberg"
	fmt.Println(s)
}

func t3() {
	var a string
	var b int
	var c bool

	fmt.Println(a, b, c)
}

func t2() {
	var x int = 1
	var y float32 = 1.00
	var z float64 = 1.0000
	var w string = "Hello world!"
	var v bool = false
	u := 2

	fmt.Println(x, y, z, w, v, u)
}

func t1() {
	fmt.Println("Hello world!")
}
