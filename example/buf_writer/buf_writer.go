package main

import (
	"bufio"
	"os"

	"github.com/ddo/go-dlog"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	logWriter := dlog.New("logWriter.Info", &dlog.Option{
		Writer: writer,
	})

	logWriter.Info("logWriter.Info 1")
	logWriter.Info("logWriter.Info 2")
	writer.Flush()

	logWriter.Info("logWriter.Info 3")
	logWriter.Info("logWriter.Info 4")
	writer.Flush()
}
