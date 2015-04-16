package main

import (
	"time"

	"github.com/ddo/go-dlog"

	"./example2"
)

func main() {
	log := dlog.New("example")
	log("some log")

	example2.Test("1234567890")
	example2.Test("example:child 2")
	example2.Test("example:child 3")
	example2.Test("example:child 4")

	time.Sleep(time.Second)

	log("some log")
}
