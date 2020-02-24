package main

import (
	"fmt"
)

func main() {

	defer do(f1(4))

	fmt.Println("hello")

}

func do(f int)  {
	fmt.Println(f)
}

func f1(x int) int {
	println(x)
	return x
}
