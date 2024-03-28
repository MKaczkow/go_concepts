# 13_testing

### General
* tools
* libraries
* test file name should end with `_test.go`
* test function name should start with `Test`
* `Fatal` and `Fatalf` functions are used to stop the test immediately, to don't waste time on the rest of the test, if some part of the test fails

> [!TIP]  
> Avoid declaring variables on package level, because they make code difficult to understand - also in tests.

* use `TestMain` func to define setup and teardown functions for test
    - it's only executed once, contrary to `setUp - tearDown` pattern in other languages
    - there can be only one `TestMain` func in a package
```go
var testTime time.Time

func TestMain(m *testing.M) {
    fmt.Println("Tests started")
    testTime = time.Now()
    exitVal := m.Run()
    fmt.Println("Tests finished in", time.Since(testTime))
    os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
    t.Log("First test", testTime)
}

func TestSecond(t *testing.T) {
    t.Log("Second test", testTime)
}
```

* ... ale jest metoda `Cleanup`, która jest wywoływana po każdym teście, jak `tearDown` w innych językach
* for suplementary data use `testdata` directory, with relative path calls

> [!TIP]  
> When testing public API leave test code in the same package as the code being tested, but create `packetname_test`, this way not-exported internals of the package won't be visible to the test code.

* `go-cmp` library is better for comparing complex data structures than stdlib `reflect.DeepEqual` [github repo link](https://github.com/google/go-cmp)
    - `cmp.Diff` returns a string with the differences between two values
    - with `cmp.Comparer` you can create custom comparing logic

* you can create custom benchmarks with starting `Benchmark` (cpt. Obvious)
* use `blackhole` variable to prevent compiler from outsmarting your benchmark
```go
var blackhole int

func BenchmarkFileLen1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result, err := FileLen("testdata/data.txt", 1)
        if err != nil {
            b.Fatal(err)
        }
        blackhole = results
    }
}
```
* there is `httptest` package for testing http handlers
* you can use `compilation flags` to run tests only in specific conditions (e.g. integration tests group only runs when some external service is available)
* `compilation flags` are defined in `// +build` comment at the beginning of the file
```go
// +build integration

package main

import "testing"

func TestIntegration(t *testing.T) {
    t.Log("Integration test")
}
```
* `race checker` is a tool to detect race conditions / issues with concurrency (cpt. Obvious)
