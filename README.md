# go-with-tests
Testing in go


## Cover profile
```
// Get test coverage
go test -cover
// Generate a cover profile and store it in wallet file (use ofter package names..)
go test -coverprofile ./wallet
// Visualize the cover profile on browser
go tool cover -html ./wallet
```

## Go benchmarking
```
go test -bench=. --benchmem
```
