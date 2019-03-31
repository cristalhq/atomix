# atomix [![GoDoc][doc-img]][doc] [![Go Report Card][reportcard-img]][reportcard]

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

[doc-img]: https://godoc.org/github.com/cristalhq/atomix?status.svg
[doc]: https://godoc.org/github.com/cristalhq/atomix
[reportcard-img]: https://goreportcard.com/badge/cristalhq/atomix
[reportcard]: https://goreportcard.com/report/cristalhq/atomix
