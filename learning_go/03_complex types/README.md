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
* ... also, `make` could be used to create dinamic slices (e.g. when you don't know the length of the slice)
```go
x := make([]int, 0, 5)
x = append(x, 1, 2, 3, 4)
```
* use `%v` to print slices
```go
fmt.Printf("%v\n", x)
```
* `...` (variadic operator) is used to unpack a slice into a variadic function
```go
for i := 0; i < currentLength; i++ {
	newSubset := append([]int{}, result[i]...)
	fmt.Printf("newSubset: %v\n", newSubset)   // %v -> printing slices
	newSubset = append(newSubset, num)
	fmt.Printf("newSubset: %v\n", newSubset)   // %v -> printing slices
	result = append(result, newSubset)
}
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
* syntax: `map[keyType]valueType`
* declaration with `var`:
```go
var nilMap map[string]int
```
* declaration with `make`:
```go
var emptyMap = make(map[string]int)
```
* declaration with `map literal`:
```go
var mapWithValues = map[string]int{"one": 1, "two": 2, "three": 3}
```
* works analogically to `slices`
* maps are not comparable, so they can only be compared to `nil`
* `comma ok` idiom:
```go
value, ok := mapWithValues["one"]
```
* enables checking if the value exists in the map (`ok` is `true` if the value exists, `false` otherwise)
* golang doesn't have `sets`, but some features can be simulated with `maps`

### Structs
* `maps` are cool, but have some drawbacks:
	- it's hard to create uniform APIs with them, because they don't have a fixed set of keys
	- all values in a map **must be** of the same type
* contary to maps, creating structs using `struct literal` and `var` keyword are equivalent (all fields will be set to their zero values, if not specified, they will be set to zero values of their types)
* most of the time, it's better to create structs, using key names, because it's more explicit and robust (e.g. if new fields will be added to the struct, the code won't break)
* `anonymous structs` - structs without a name, used for temporary data (no `type` keyword), e.g. for testing purposes
* `maps` and `slices` are not comparable, so struct containing fields with these types won't be comparable as well
