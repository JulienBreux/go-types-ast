//go:build exclude

package example

var x int

func _() {
	x := x
	_ = x
}
