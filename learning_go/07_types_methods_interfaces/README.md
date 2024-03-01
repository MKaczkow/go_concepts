# 07_types_methods_interfaces

### Intro
* `Golang` is statically typed language
* avoid inheritance, use composition instead
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
