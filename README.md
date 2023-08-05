# atomix

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

Better `sync/atomic` package for Go.

## Rationale

To make life easier with `sync/atomic` this package provide wrappers and helper functions, so all the atomic data and operations are easily visible.

## Note

For the better `sync` package see [synx](https://github.com/cristalhq/synx).

## Features

* All primitive types and more
  * `int8`, `int16`, `int32`, `int64`
  * `uint8`, `uint16`, uint32`, `uint64`
  * `float32`, `float64`, `complex64` and `int`
  * `string`, `interface{}` and `error`
  * `uintptr`, `unsafe.Pointer`
  * `time.Time`, `time.Duration`
* Zero cpu and memory overhead in almost all cases.
* Useful helpers.

See [docs][pkg-url].

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

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/atomix/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/atomix/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/atomix
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/atomix
[version-img]: https://img.shields.io/github/v/release/cristalhq/acmd
[version-url]: https://github.com/cristalhq/acmd/releases
