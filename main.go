package main

import "fmt"

func main() {

	res := add(3, 4)
	fmt.Println("The result is ", res)
}

func add(one int, two int) int {

	return one + two
}
