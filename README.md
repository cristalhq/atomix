# atomix

Simple and easy wrappers for Go `sync/atomic` package.

## Install

```
go get github.com/cristalhq/atomix
```

## Example

```go
var a atomix.Int32
a.Store(1335)
a.Add(2)
b := a.Load()
```

## License

[MIT License](LICENSE).
