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
* type declarations doesn't mean inheritance
* no `enum` (enumarating type) in Golang, but can be simulated with `const` and `iota`
    ```go
    package main

    import "fmt"

    type Suit int

    const (
        Spades Suit = iota
        Hearts
        Diamonds
        Clubs
    )

    func main() {
        var cardSuit Suit = Spades
        fmt.Println(cardSuit)   // 0
    }
    ```
* embedded fields as a way of simulating composition
* no `dynamic dispatch` in Go for concrete types 

### Methods
* `methods` are functions that are associated with a type (they have `receiver`, kinda like `self` in Python)

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
* programming with `nil` value in mind
    - expecting `nil` receiver
    - most languages doesn't allow executing methods on `nil` receiver
    - Golang allows it
    - if this is `value receiver` method -> `panic`
    - if this is `pointer receiver` method -> `nil` is passed to the method (it may work, if written 'with `nil` in mind')
* methods are also functions (kinda like static)
    ```go
    type Adder struct {
        start int
    }

    func (a Adder) AddTo(val int) int {
        return a.start + val
    }
    ```
    - normal way
    ```go
    myAdder := Adder{start: 10}
    fmt.Println(myAdder.AddTo(5))   // 15
    ```
    - `method value`
    ```go
    f1 := myAdder.AddTo
    fmt.Println(f1(5))   // 15
    ```
    - `method expression`
    ```go
    f2 := Adder.AddTo
    fmt.Println(f2(myAdder, 15))   // 25
    ```

### Interfaces
* the only abstract type in Golang is `interface`
```go
type Stringer interface {
    String() string
}
```
* no need to declare that a `concrete type` implements an `interface`, it's done implicitly
* this combines pros of dynamic languages and static languages
    - `duck typing` (like in Python)
    - `interfaces` (like in Java)

* in Golang, we have interface, but only client function 'knows' about it 
* thus, we can easly change the implementation of the interface (like in dynamic languages)
* but, the compiler will check if the new implementation is compatible with the interface (like in static languages)
```go

type LogicProvider struct {}

func (lp LogicProvider) Process(data string) string {
    return "processed: " + data
    // program login
}

type Logic interface {
    Process(string) string
}

type Client struct {
    L Logic
}

func (c Client) Program(data string) string {
    // get data
    return c.L.Process(data)
}

main() {
    c := Client{
        L: LogicProvider{},
    }
    c.Program()
}
```
* interfaces can be shared

> [!TIP]  
> If there is interface in stdlib, which fits your needs, use it. You may adhere to the decorator pattern - create a factory function, which accept interface instance and return different type, implementing the interface.

* interfaces can be embedded (eg. `io.ReadCloser` is an interface that embeds `io.Reader` and `io.Closer`)

> [!TIP]  
> Accept interfaces, return structs.

* because of lack of generics, `interface{}` is used to represent any type

### Miscellaneus
* ... but how to check if `interface` type variable is of specific `concrete` type?
* `type assertion` is used to extract the value from an interface
```go

type MyInt int

func main() {
    var i interface{}
    var mine MyInt = 42
    i = mine
    i2 := i.(MyInt)   // this is type assertion
    fmt.Println(is + 1)
}

```
* `type conversions` change type, while `type assertions` extracts it
* other way is to use `type switch`
* `DIP` (Dependency Inversion Principle) example in Golang
```go

func LogOutpus(message string) {
    fmt.Println(message)
}

type SimpleDataStore struct {
    userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(id string) (string, bool) {
    name, ok := sds.userData[userID]
    return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
    return SimpleDataStore{
        userData: map[string]string{
            "1": "Fred",
            "2": "Mary",
            "3": "Pat",
        },
    }
}



```