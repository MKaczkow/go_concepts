# Go Concepts
Repo for basic tutorial-based Golang study  

---

# dubious shadowing
* `shadowing` - variable with the same name as another variable in the same scope, which hides the other variable
* `static analysis` sometimes helps
* ... but `control flow graphs` are needed other times

# references
- [cool blogpost](https://jakebailey.dev/posts/go-shadowing/)
- [go/analysis docs](https://pkg.go.dev/golang.org/x/tools/go/analysis)
- [shadow analyzer docs](https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/shadow)
- [typescript-go issue on shadowing](https://github.com/microsoft/typescript-go/pull/365)
- [typescript issue on control flow analysis](https://github.com/microsoft/TypeScript/issues/9998)