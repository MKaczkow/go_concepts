# 12_context

### Intro
* context is just an instance of something which implements `context.Context` interface
* in the beginning of the program, `context.Background` may be used, to create blank `context.Context` instance, this is an exception to the rule that functions should return concrete types

> [!TIP]  
> Context should be used to pass values *deeper* into the call stack, but not the other way around, because it's treated as immutable and wrapped each time it's passed to another function, with added values

### Cancellation
* use `context.WithCancel` to create new context, which can be canceled

### Timeout
* use `context.WithTimeout` to create new context, which will be canceled after specified time
* if context still active when `Done` is called, it will return `nil`, otherwise it will return one of the sentinel errors: 
    - `context.Canceled` 
    - `context.DeadlineExceeded`

#### How to manage server workload?
* restrict number of concurrent requests => restrict number of `goroutines`
* restrict number of awaiting requests => use `buffered channels`
* restrict time of processing a request
* restrict amount of resources used by server while processing a request

#### General pattern to manage long running stuff in your code
```go
func longRunningStuff(ctx context.Context, data string) (string, error) {
    type wrapper struct {
        result string
        err error
    }
    ch := make(chan wrapper, 1)   // buffered channel with size 1 will allow goprocedure to finish, even if buffered value will never be read, because of context cancelation
    go func() {
        // do long running stuff
        result, err := longRunningStuff(ctx, data)
        ch <- wrapper{result, err}
    }()
    select {
        case data := <-ch:
            return data.result, data.err
        case <- ctx.Done():
            return "", ctx.Err()
    }
}
```

### Metadata
* context may also serve one more purpose, which is to pass metadata between functions 
* ... which really shouldn't be done (explicite arguments are better)

### Smth with GUID
* example app in this [Github repo](https://github.com/learning-go-book/context_values)