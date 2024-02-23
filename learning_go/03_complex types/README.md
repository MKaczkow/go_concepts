# 03_complex_types

* `slices` are more flexible than `arrays`
* `var []int{1, 2, 3}` - declares a slice
* `var [...]int{1, 2, 3}` - declares an array
* `slices` may only be compared to `nil` (otherwise use `reflect.DeepEqual` from `reflect` package or get compilation error)
* `unicode` uses 4 bytes (32 bits) for each character, so 'default' way to represent data should be UTF-32, right?
* but it's very inefficient, so `utf-8` is used instead, which enables to represent most of characters with 1 byte (8 bits) 

### Creating slices with `make`, `var` and `literal`
* use `literal` when you roughly know initial values or when you are pretty sure they won't change
```go
x := []int{1, 2, 3}
```
* use `make` when you know the length of the slice but not the values
```go
x := make([]int, 3)
```
* use `var` when it's possible that the slice length will be zero and won't change
```go
var x []int
```

### Slice expressions
* don't create a new slice, just a reference to the original slice
* not-used capacity of original slice is shared with it's 'derivative' slices
* example:
```go
package main

import "fmt"

func main() {
	x := make([]int, 0, 5)   // create a slice with length 0 and capacity 5
	x = append(x, 1, 2, 3, 4)   // add 4 elements to the slice
	y := x[:2]   // y is [1, 2, CAP, CAP, CAP]
	z := x[2:]  // z is [3, 4]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)   // y is [1, 2, 30, 40, 50]
	x = append(x, 60)   // x is [1, 2, 30, 40, 60]
	z = append(z, 70)   // z is [30, 40, 70] (starts with [2], overwrites x and y (XDD))
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
```
* output:
```go
5 5 3
x: [1 2 30 40 70]
y: [1 2 30 40 70]
z: [30 40 70]
```
* why did this happen?
* (should be)`x` has a capacity of 5, `y` and `z` are slices of `x` and share the same capacity, but (as is) `z` capacity is smaller

> [!WARNING]  
> Need to be very careful while creating slices-out-of-slices. This may cause weird issues, as it's copy-by-reference, not by value. Use `full slice expressions` (specifying *last posistion in source slice capacity to be available for child slice*) to create a new slice with the same values as the original slice.

### Maps and hash maps
TBD