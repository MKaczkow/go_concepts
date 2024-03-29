# 15_parametric_types

### General
* powerful, diligent description of proposal could be found [here](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
* Golang is statically typed => types of variables and arguments are checked on compile time 
* `generics` are a way to write functions and data structures that can work with any type (or most of them)
* this may be simulated using `closures` as shown [here](https://medium.com/capital-one-tech/closures-are-the-generics-for-go-cb32021fb5b5)
* problem (choose 1):
    - `slow programmers`
    - `slow compilers`
    - `slow execution times`
* you need type safety (checking on compile time)
* go `operator overloading`
