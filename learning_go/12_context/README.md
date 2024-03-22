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

#### How to manage server workload?
* restrict number of concurrent requests => restrict number of `goroutines`
* restrict number of awaiting requests => use `buffered channels`
* restrict time of processing a request
* restrict amount of resources used by server while processing a request
