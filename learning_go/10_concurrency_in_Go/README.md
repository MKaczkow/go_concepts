# 10_concurrency_in_Go

### Intro
* concurrency in Go doesn't use locks
* it's based on `communicating sequential processes` (CSP) model [reference paper](https://dl.acm.org/doi/pdf/10.1145/359576.359585)
* interesting book `The Art of Concurrency` by Clay Breshears

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
    * `Add` method to add number of goroutines to wait for
    * `Done` method to signal that goroutine has finished
    * `Wait` method to block until all goroutines have finished
* there is also `ErrGroup`
* when to use `buffered`, when to use `unbuffered` channels?
* `buffered` - useful when you know how many goroutines are created, want to reduce their number or reduce number of awaiting operations
```go
func processChannel(ch chan int) [] int {
    const conc = 10
    results := make(chan int, conc)
    for i := 0; i < conc; i++ {
        go func() {
            v := <- ch
            results <- process(v)
        }()
    }
    var out []int
    for i := 0; i <conc; i++ {
        out = append(out, <-results)
    }
    return out
}
```
* `backpressure` mechanism - basically, there is a buffered channel with *chips*, befere doing thing, goroutine tries to obtain a chip, if it succeeds, it can do the action, if it fails there is error like 'no more system resources'
* `sync.Once` is for lazy initialization
```go

type SlowComplicatedParser interface {
    Parse(string) (string)
}
var parser SlowComplicatedParser
var once sync.Once

func Parse(dataToParse string) string {
    once.Do(func() {
        parser = initParser()
    })
    return parser.Parse(dataToParse)
}

func initParser() SlowComplicatedParser {
    // do some heavy initialization
}
```

### Select 
* `select` statement is used to wait for multiple channels to send or receive values
* like `switch`, but chooses random satisfied case
* but deadlock is still possible xd
* `for-select` loop
```go
for {
    select {
    case v := <-ch1:
        // do something with v
    case v := <-ch2:
        // do something with v
    }
}
```
* `done channel pattern` ...
```go
for {
    select {
    case v := <-ch1:
        // do something with v
    case v := <-ch2:
        // do something with v
    case <-done:
        return
    }
}
```
* ...but it's better to not use `default`

### Best practices
* `API`s shouldn't be designed to use concurrency, because it's implementation detail, which should be hidden
* so, *don't expose anything concurrency-related in your API* seems to be a good rule of thumb
* avoid `goroutine leaks` 
* it's possible to set time limit to op using `time.After` - this is called `timeout idiom`
```go
func timeLimit() (int, error) {
    var result int
    var err error
    done := make(chan struct{})
    go func() {
        result, err = longRunningOp()
        close(done)
    }()
    select {
        case <-done:
            return result, err
        case <-time.After(2 * time.Second):
            return 0, errors.New("timeout")
    }
}
```
