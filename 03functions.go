package main

import "fmt"

func main() {
	welcome()
	addition(5, 3)
	result := multiply(4, 10)
	fmt.Println(result)
}

func addition(i1, i2 int) {
	panic("unimplemented")
}

func multiply(i1 int, i2 int) {
	return i1 * i2
}
