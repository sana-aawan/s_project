package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverse_num(45))

}

func reverse_num(n int) int {
	//rem := 0
	new_int := 0
	for n > 0 {
		remainder := n % 10
		new_int *= 10
		new_int += remainder
		n /= 10
	}
	return new_int
}