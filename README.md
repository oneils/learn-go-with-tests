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

## Using Mutex or Channels

A `Mutex` is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
`sync.WaitGroup` which is a convenient way of synchronising concurrent processes

- Use channels when passing ownership of data 
- Use mutexes for managing state

## go vet

go vet

Remember to use `go vet` in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.