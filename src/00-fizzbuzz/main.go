package main

import "fmt"

const f = 3
const b = 5

func main() {

	for i := 0; i <= 100; i++ {
		// if i is devisable by f
		if i%f == 0 && i%b != 0 {
			fmt.Println("fizz")
		}
		// if i is devisable by b
		if i%b == 0 && i%f != 0 {
			fmt.Println("buzz")
		}
		// if i is devisable by both f and b
		if i%f == 0 && i%b == 0 {
			fmt.Println("fizzbuzz")
		}
		// if i is devisable by both f and b
		if i%f != 0 && i%b != 0 {
			fmt.Println(i)
		}
	}
}
