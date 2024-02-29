# 05_functions

### Intro
* functions are `first-class citizens`
* Go is strongly typed language, thus arguments and return values must have a type
* no named parameters (`kwargs`) and no default values
* `variadic` functions are possible (they take a variable number of arguments)
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```
* multiple return values are possible
```go
func divide(a, b int) (int, int) {
    return a / b, a % b
}
```
* `error` should be the last or only return value
* named return values are supported
```go
func divide(a, b int) (result int, remainder int, err error) {
    if b == 0 {
        err = errors.New("cannot divide by zero")
        return
    }
    return a / b, a % b, nil 
}
``` 
* `func` is a keyword, but also a type, so it can be used as a value
```go
type opFuncType func(int, int) int   // any function that takes two ints and returns an int
```
* `anonymous functions` are basically useful in two cases
    - using `defer`
    - using `goroutines`
* `closures`
* `defer` is a keyword that schedules a function call to be run after the function completes, useful for cleanup
* `call by value` means that the function gets a copy of the value, not the original value (always!)
