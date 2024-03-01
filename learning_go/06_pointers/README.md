# 06_pointers

### Intro
* `garbage collector` is present
* `nil` can't be converted to a numerical (or any) value
* `pointer arithmetic` is not allowed
* `*` is dereference operator (changes reference-type variable into it's value)
* `&` is reference operator (changes value-type variable into it's reference)
```go
package main

import "fmt"

func main() {
    var a int = 42
    var b *int = &a
    fmt.Println(a, *b)   // 42 42
    a = 27
    fmt.Println(a, *b)   // 27 27
}
```
* `new` is used to create a pointer to a value-type variable, but it's not used rarely

### Efficiency
* passing a pointer to a function is more efficient than passing a value, because pointer has a fixed size, while value can be large
* but returning a pointer from a function 

### Maps vs slices
* map is a reference type, slice is a value type
* thus, maps should be avoided as arguments, especially in large, public APIs
* `slices` (the only practically useful linear data structure in Go) are often passed, as functions' arguments, but should not be modified inside function body

### Heap vs stack
* `heap` is used for dynamic memory allocation, and is managed by garbage collector
* `stack` is used for static memory allocation, and is managed by programmer
* each `goroutine` has it's own stack, all managed by `Go runtime`
* `mechanical sympathy` is understanding how the hardware works, and how the software interacts with it, to write more efficient code
