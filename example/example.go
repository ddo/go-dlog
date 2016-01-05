package main

import (
	"bufio"
	"os"

	"github.com/ddo/go-dlog"
)

func main() {
	logDefault := dlog.New("logDefault", nil)
	logStdout := dlog.New("logStdout", os.Stdout)
	logStderr := dlog.New("logStderr", os.Stderr)

	writer := bufio.NewWriter(os.Stdout)
	logWriter := dlog.New("logWriter", writer)

	logDefault("logDefault 1")
	logDefault("logDefault 2")

	logStdout("logStdout 1")
	logStdout("logStdout 2")

	logStderr("logStderr 1")
	logStderr("logStderr 2")

	logWriter("logWriter 1")
	logWriter("logWriter 2")
	writer.Flush()
}
