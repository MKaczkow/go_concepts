# 07_types_methods_interfaces

### Intro
* `Golang` is statically typed language
* avoid inheritance, use composition instead

### Types
* 'Person' type with base type defined by the following struct literal
```go
package main

type Person struct {
    FirstName string
    LastName  string
    Age       int
}
```
* also possible with simple types
```go
type Score int
type Converter func(string) Score   // 'Converter' is a function type, function that takes a string and returns a 'Score' 
type TeamScores map[string]Score   // 'TeamScores' is a map type, map from string to 'Score'
```
* `abstract types` defines 'what is to be done' and **not** 'how to do it'
* `concrete types` defines 'what is to be done' and 'how to do it'
* in Golang all types are either `abstract` or `concrete`, no way of creating hybrid types (like interface with default methods in Java)

### Methods
* `methods` are functions that are associated with a type (they have `receiver`, kinda like `self` in Python)
* 
```go
package main

import "fmt"

type Person struct {
    FirstName string
    LastName  string
    Age       int

    func (p Person) FullName() string {   // p Person is a receiver
        return p.FirstName + " " + p.LastName
    }

}
```
* can't overload methods in Golang, but can have multiple methods with the same name if they operate on different types
* `pointer arguments` are used to indicate that the method will modify the argument
* ... but most of the time it's better to use `value arguments`
* this is also true for `pointer receivers` and `value receivers`
    - if method modifies the receiver -> `pointer receiver` must be used
    - if method must handle `nil` instance -> `pointer receiver` must be used
    - if method doesn't modify the receiver -> `value receiver` can be used
* programming with `nil` value in mind (TBD)