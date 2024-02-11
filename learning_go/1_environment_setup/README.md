# 1_environment_setup

* `go run` - treat program as script (compile, run and delete binary)
* `go build` - compile program
* `hey` - load testing tool [github link](https://github.com/rakyll/hey)
* `go vet` - check for suspicious constructs, different than `golint` (checks for style) and `gofmt` (formatting), so the correct order of usage is: 
    1. `gofmt`  
    2. `go vet`  
    3. `golint`
* `go vet` uses some fancy heuristics to detect issues, which may have been missed by compiler - not all of of are genuine problems (FP exists) [docs link](https://golang.org/cmd/vet/)
* [golang backward compatibility blogpost](https://go.dev/blog/compat) and [golang backward compatibility doc](https://go.dev/doc/go1compat) describes promise of backward compatibility in golang made by golang team (the boring programming language), there are some interesting caveats, like problem with too-precise `time.Now` and mismatch between go and 'rest of the world' on IP addresses with leading zeros. Also compatibility is guaranteed at source level, not binary level, so recompile after version update
