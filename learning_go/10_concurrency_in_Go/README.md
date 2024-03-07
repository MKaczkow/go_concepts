# 10_concurrency_in_Go

### Intro
* concurrency in Go doesn't use locks
* it's based on `communicating sequential processes` (CSP) model [reference paper](https://dl.acm.org/doi/pdf/10.1145/359576.359585)
* interesting book `The Art of Concurrency` by Clay Breshears
* 

### Goroutines 
* basic building block of concurrency in Go
* goroutine is a process managed by Go runtime
* Go runtime has a scheduler, so goroutines don't really use OS scheduler

### Channels
* built-in type for communication between goroutines
* created with `make` function
```go 
ch := make(chan int)
```
* similar to maps, `channels` are reference types
* operator `<-` is used to send and receive messages
```go
a := <-ch // receive from ch and assign to a
ch <- b // send b to ch
```
* it's good practice to define channels as `read-only` or `write-only` as much as possible
```go 
func read(ch <-chan int) {
    // read-only channel
    // function body
}

func write(ch chan<- int) {
    // write-only channel
    // function body
}
```
* by default channels are `unbuffered`, meaning that they can only hold one value at a time
* `buffered` channels can hold more than one value
```go
ch := make(chan int, 10) // buffered channel with capacity of 10
```
* built-in functions `len` and `cap` can be used to get length and capacity of the channel
* `close` function is used to close the channel
    - after closing the channel, it's not possible to send any more values to it (`panic` will occur)
    - after closing the channel, it's still possible to receive values from it (if there are still valus in buffer they're returned, if not `zero` value, for given channel type, is returned)
* use `comma, ok` idiom to check if received `zero` value is actually a value from the channel or a channel has been closed
```go
v, ok := <-ch
```
* use `sync.WaitGroup` to wait for all goroutines to finish and avoid panic

### Best practices
