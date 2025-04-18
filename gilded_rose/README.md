# Go Concepts
Repo for basic tutorial-based Golang study  

---

### gilded rose
* this is a 'famous' exercise, which is supposed to be a refactoring kata
* the task is to refactor the code in a way, that it's easier to understand and maintain
* code has been copied from original repo
* task: add some feature -> `conjured` item

### makefile
* `make run NUM_DAYS=21`
* `make test`
* `make test-coverage`

### original readme
- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```

### resources
* [original repo, golang subdir](https://github.com/emilybache/GildedRose-Refactoring-Kata/tree/e2abba77cb5a395702f237e428b639f2129b1f07/go)
* [original requirements, in PL](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements_pl.md)
* [original requirements, in EN](https://github.com/emilybache/GildedRose-Refactoring-Kata/blob/main/GildedRoseRequirements.md)
* [dev.to article about the task](https://dev.to/lomig/a-walk-through-the-gilded-rose-kata-pt-1-do-not-break-anything-23b1)
* [another dev.to article](https://dev.to/alexandreruban/lessons-from-the-gilded-rose-refactoring-kata-pgh)
* [pretty cool blog article](https://blog.lunarlogic.com/2015/what-ive-learned-by-doing-the-gilded-rose-kata-4-refactoring-tips/)