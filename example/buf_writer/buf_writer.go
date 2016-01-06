package main

import (
	"bufio"
	"os"

	"github.com/ddo/go-dlog"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	logWriter := dlog.New(&dlog.Option{
		Name:   "logWriter",
		Writer: writer,
	})

	logWriter("logWriter 1")
	logWriter("logWriter 2")
	writer.Flush()

	logWriter("logWriter 3")
	logWriter("logWriter 4")
	writer.Flush()
}
