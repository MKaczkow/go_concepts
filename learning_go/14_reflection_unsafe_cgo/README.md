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
* `reflect.New` is reflective version of `new` operator
* there are also equivalents of `make` 
* it's possible to create custom marshalling and unmarshalling functions using `reflect` package and reflection mechanism

### Unsafe
* truly unsafe
* weird
* 1 type -> `Pointer`
* 3 functions
    - `Sizeof` -> returns size of the variable
    - `Offsetof` -> returns offset of the field in the struct
    - `Alignof` -> returns alignment of the field in the struct
* based on [this empirical study from arxiv](https://arxiv.org/pdf/2006.09973.pdf) some colorful stuff

> [!TIP]
> Correct way to convert `float64` to `uint64` using `unsafe` package (conversion must appear in the same line s operations, otherwise the pointer could be collected by garbage collector before dereferencing) 
> ```go
> import unsafe
> 
> func Float64(f float64) uint64 {
>    return *(*uint64)(unsafe.Pointer(&f))
> }
> ```

> [!CAUTION]
> Incorrect way, contrary to explained above
> ```go
> import unsafe
> 
> u := uintptr(p)
> p = unsafe.Pointer(u + offset)
> ```


### Cgo
