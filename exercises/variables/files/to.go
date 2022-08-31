//go:build exclude

package main

import "fmt"

var y int

func main() {
	if y == 0 {
		y = 123
	}

	fmt.Printf("%d\n", y)
}

func _() {
	x := y
	_ = x
}
