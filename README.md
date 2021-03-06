# atomix

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

Simple and easy wrappers for Go `sync/atomic` package.

## Features

* All primitive types and more
  * int8, int16, int32, int64
  * uint8, uint16, uint32, uint64
  * float32, float64, complex64
  * string, interface{} and error
  * uintptr, unsafe pointer
  * time.Time, time.Duration
* Zero cpu and memory overhead.
* Additional methods.

## Install

```
go get github.com/cristalhq/atomix
```

## Example

```go
var a atomix.Int32
a.Store(1335)
a.Add(2)
b := a.Load() // 1337
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/atomix/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/atomix/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/atomix
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/atomix
[reportcard-img]: https://goreportcard.com/badge/cristalhq/atomix
[reportcard-url]: https://goreportcard.com/report/cristalhq/atomix
[coverage-img]: https://codecov.io/gh/cristalhq/atomix/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/atomix
