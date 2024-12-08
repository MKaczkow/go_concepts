# Go Concepts
Repo for basic tutorial-based Golang study  

---

### todo
- [ ] running and testing pipeline
- [ ] parsing 
  - [ ] or
  - [ ] repeat 
  - [ ] repeatspecified
- [ ] compiling
- [ ] matching engine

# regex engine
* based on [this blogpost](https://rhaeguard.github.io/posts/regex/)

## run
* `go test ./...`

## notes
* 3 stages:
  * `parse` - create tokens from string
  * `build state machine(compile)` - create state machine from tokens
  * `match` - match string with state machine
* `NFA` - non-deterministic finite automaton
* `DFA` - deterministic finite automaton

## references
* [wikipedia on NFA](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton)
