package main

import "fmt"

func add(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtract(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func multiply(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}

func main() {
	sum := add(1, 2, 3)
	sub := subtract(5, 10)
	mult := multiply(10, 10)
	div := divide(20)
	fmt.Println(sum, sub, mult, div)
}
