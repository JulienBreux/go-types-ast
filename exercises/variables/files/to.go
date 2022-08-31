//go:build exclude

package example

var y int

func _() {
	x := y
	_ = x
}
