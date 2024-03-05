# 08_errors

### Intro
* `panic` and `recover`
* controversial topic
* `sentinel`-type errors
* `nil` is zero-value for any type
* 2 ways to create errors, based of string:
    - `errors.New("message")`
    - `fmt.Errorf("message")`

#### Sentinel errors
* should be treated as `read-only` values
* shouldn't be modified or created

#### Errors are values
* `error` is an interface
* thus, you can implement your own error
* for interface to be equal to `nil` (`== nil`), both base type and value should be `nil`
* most common way is to return `nil` as error's value if no error occurred

#### Wrapping errors
* sequence of errors, wrapped is called `error chain`
* `errors.Is` and `errors.As` are used to work with error chains

#### Panic and recover
* shouldn't be overused
