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
