package main

import "fmt"

func main() {
	fmt.Println("hi")

	a := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	b := a[1:3]
	fmt.Println(b)
}
