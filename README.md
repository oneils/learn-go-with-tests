This is a project for learning Go by TDD via https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers

## Cross compile

```bash
env GOOS=linux GOARCH=arm go build
```
[Go (Golang) GOOS and GOARCH](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63)


## Benchmarks

To run the benchmarks do `go test -bench=.`

## Tests coverage

Run
```go
go test -cover
```