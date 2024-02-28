# 04_blocks_overloading_and_control_flow

* any place, where there is a declaration of something, is a `block`

### Shadowing
* if a variable is declared in a block, it's visible only in that block
```go
package main

import "fmt"

func main() {
    x := 10
    {
        x := 20
        fmt.Println(x)   // prints 20
    }
    fmt.Println(x)   // prints 10
}
```
* interestin `shadow` linter

### For loop
* 4 types
    * `for` - like `while` in other languages
    * `for` with `init` and `post` statements
    * `for` with `condition` only
    * `for` with `range` - used to iterate over a collection
```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }

    for {
        fmt.Println(i)
        i++
        if i == 5 {
            break
        }
    }

    s := []int{1, 2, 3}
    for k, v := range s {
        fmt.Println(k, v)
    }
}
```
* while iterating over `maps`, with `for-range`, the order of the elements is not guaranteed
```go
package main

import "fmt"

func main() {
    m := map[string]int{
        "a": 1,
        "b": 2,
        "c": 3,
    }

    for k, v := range m {
        fmt.Println(k, v)
    }
}
```
* this is security measure againt `hash DoS` attacks (if data is stored in a `map` on a server, the attacker can send prepared data, so that it will be stored in the same bucket)
* `for` loop doesn't read multi-bytes characters correctly
```go
package main

import "fmt"

func main() {
    s := "Hello, 世界"
    for k, v := range s {
        fmt.Println(k, v)
    }
}
```
### Switch statement
* `expression switch` vs `type switch`
    - `expression switch` - `case` values are expressions
    - `type switch` - `case` values are types
* `case` clauses don't fall through, thus no need for `break` statement
* `fallthrough` statement - forces the next `case` to be executed (but why would you want to do that? - think simpler)
* `blank switch` statements 
