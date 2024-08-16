package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	const pi = 3.14159
	const projectName = "My First Go project!"

	var a int = 10
	var b float64 = 3.14
	var name string = "David"
	var isInitialised bool = true

	i := 100
	j := 2.8234

	var david Person
	david.name = "David"

	var m []int = []int{1, 2, 3}
	n := []string{"Go", "Python", "Java"}

	fmt.Printf("a: %d, type: %T \n", a, a)
	fmt.Printf("b: %f, type: %T \n", b, b)
	fmt.Printf("name: %s, type: %T \n", name, name)
	fmt.Printf("isInitialised: %t, type: %T \n", isInitialised, isInitialised)
	fmt.Printf("i: %d, type: %T \n", i, i)
	fmt.Printf("j: %f, type: %T \n", j, j)
	fmt.Printf("j: %f, type: %T \n", j, j)
	fmt.Printf("david: %+v, type: %T \n", david, david)
	fmt.Printf("j: %v, type: %T \n", m, m)
	fmt.Printf("j: %v, type: %T \n", n, n)
}
