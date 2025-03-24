package main

import (
	"fmt"
)

type G[T any] struct {
	Data []T
}

type D = G[int] // <- type alias or type definition both can trigger this bug.

func main() {
	g := D{
		Data: make([]int, 0), // goland will report a type error here.
	}
	fmt.Println(g)
}
