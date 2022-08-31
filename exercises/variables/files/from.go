//go:build exclude

package main

import "fmt"

var x int

func main() {
	if x == 0 {
		x = 123
	}

	fmt.Printf("%d\n", x)

}

func _() {
	x := x
	_ = x
}
