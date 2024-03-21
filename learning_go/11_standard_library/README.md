# 11_standard_library

### Intro
* official documentation [here](https://golang.org/pkg/)
* example elements of standard library:
    - I/O ops
    - time
    - JSON
    - HTTP

### I/O ops
* `io.Reader` and `io.Writer` interfaces
* `io.EOF` error, which is not true error
* interfaces are safe, type-wise (this weird combination of `duck typing` and `not-duck typing`)
* `io.Copy`, as name implies, copies data between `io.Reader` and `io.Writer` 
* `io.Closer` is implemented by `types` like `os.File`, in cases, where you need to clean up something
* `io.Seeker` enables accesing data in random place of resource
* cool pattern `ioutil`, shows how to adding method to type, which is not really implementing them, to adhere to some interface (kinda like mocking)

```go
type nopCloser struct {
    io.Reader
}

func (nopCloser) Close() error { return nil}

func NopCloser(r io.Reader) io.ReadCloser {
    return nopCloser{r}
}
```

* but it breaks general rule of `don't return interfaces in functions`

### Time
* `time.Duration`
    * `time.Hour`
    * `time.Minute`
    * ...
    *  nanosec is the smallest one
* `time.Time`
    * timezone info is included, so don't compare using `==`, but using `Equals`
* a little weird format

```go
t, err := time.Parse("2006-02-01 15:04:05 -0700", "2016-13-03 00:00:00 +0000")
if err != nil {
    return err
}
fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
```
* now, time in Go guarantees monotonicity (uses monotonic clock, instead of real time)
* `time.Now()` used to rely on real time, which caused bug described on [cloudflare blog](https://blog.cloudflare.com/how-and-why-the-leap-second-affected-cloudflare-dns/)

### JSON

### HTTP
