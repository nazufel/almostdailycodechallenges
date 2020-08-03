package main

import "fmt"

const f = 3
const b = 5

func main() {

	for i := 0; i <= 100; i++ {
		r := fizzBuzz(i)
		if len(r) != 0 {
			fmt.Println(r)
		}
		if len(r) == 0 {
			fmt.Println(i)
		}

	}

}

func fizzBuzz(c int) string {

	var res string

	// if i is devisable by f
	if c%f == 0 && c%b != 0 {
		res = "fizz"
	}
	// if i is devisable by b
	if c%b == 0 && c%f != 0 {
		res = "buzz"
	}
	// if i is devisable by both f and b
	if c%f == 0 && c%b == 0 {
		res = "fizzbuzz"
	}

	return res
}
