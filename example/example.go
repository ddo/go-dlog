package main

import (
	"time"

	"github.com/ddo/go-dlog"

	"./example2"
)

func main() {
	log := dlog.New("example")
	log("some log")

	example2.Test("example 2")
	example2.Test("example 3")
	example2.Test("example 4")
	example2.Test("example 5")

	time.Sleep(time.Second)

	log("some log")
}
