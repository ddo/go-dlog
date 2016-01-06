package main

import (
	"os"

	"github.com/ddo/go-dlog"
)

func main() {
	file, err := os.Create("/tmp/dlog_test.txt")

	if err != nil {
		panic(err)
	}

	logWriter := dlog.New(&dlog.Option{
		Name:   "logWriter",
		Writer: file,
	})

	logWriter("logWriter 1")
	logWriter("logWriter 2")

	file.Sync()

	// check /tmp/dlog_test.txt
}
