# go-dlog [![Doc][godoc-img]][godoc-url]
:pager: conditional logging for Golang libraries &amp; applications

![screenshot][screenshot]

> inspired by TJ [debug](https://github.com/visionmedia/debug)

## installation

```sh
go get github.com/ddo/go-dlog
```

## usage

```go
log := dlog.New("example")
log("some log")
```

```sh
DLOG=* go run example.go
```

set ***DLOG*** environment to any to enable the logging

[godoc-img]: https://img.shields.io/badge/godoc-Reference-brightgreen.svg?style=flat-square
[godoc-url]: https://godoc.org/github.com/ddo/go-dlog
[screenshot]: http://i.imgur.com/b7OcA3R.png
