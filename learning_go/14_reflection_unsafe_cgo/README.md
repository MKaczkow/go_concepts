# 14_reflection_unsafe_cgo

### Intro
* `reflection` -> when you can't know the type of the variable at `compilation` time
* `unsafe` -> when you need to bypass Go's type system
* `cgo` -> when you need to call C code from Go

### Reflection
* just use Python XD
* `reflect` package
* often used *on the edge* between program and real world
* difference between `Kind` and `Type`
    - `Kind` is the basic type of the variable
    - `Type` is the full type of the variable
    - when you define `Foo` struct, it has:
        - `Kind` -> `reflect.Struct`
        - `Type` -> `packetname.Foo`

### Unsafe

### Cgo