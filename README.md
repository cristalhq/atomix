# atomix

[![Build Status][travis-img]][travis-url]
[![GoDoc][doc-img]][doc-url]
[![Go Report Card][reportcard-img]][reportcard-url]
[![Go Report Card][coverage-img]][coverage-url]

Simple and easy wrappers for Go `sync/atomic` package.

## Features

* All primitive types.
* Zero performance overhead.
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
b := a.Load()
```

## Documentation

See [these docs](https://godoc.org/github.com/cristalhq/atomix).

## License

[MIT License](LICENSE).

[travis-img]: https://travis-ci.org/cristalhq/atomix.svg?branch=master
[travis-url]: https://travis-ci.org/cristalhq/atomix
[doc-img]: https://godoc.org/github.com/cristalhq/atomix?status.svg
[doc-url]: https://godoc.org/github.com/cristalhq/atomix
[reportcard-img]: https://goreportcard.com/badge/cristalhq/atomix
[reportcard-url]: https://goreportcard.com/report/cristalhq/atomix
[coverage-img]: https://coveralls.io/repos/github/cristalhq/atomix/badge.svg?branch=master
[coverage-url]: https://coveralls.io/github/cristalhq/atomix?branch=master
