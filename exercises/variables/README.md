# Variables renaming exercise

The main purpose of this exercise is to check how to rename a variable already in use.

## Instruction

Rename the var `x` to var `y`.

## Help

1. https://pkg.go.dev/go/types#example-Info

## Structure

`from.go`:

```go
package example

var x int

func _() {
	x := x
	_ = x
}
```

 `to.go`:

 ```go
package example

var y int

func _() {
	x := y
	_ = x
}
 ```

 ## Test

    go test -timeout 30s -v -count=1 .

## To Do

- [ ] Check variable usage
