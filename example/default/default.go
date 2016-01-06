package main

import (
	"os"

	"github.com/ddo/go-dlog"
)

// if nil writer -> os.Stdout
func main() {
	logDefault := dlog.New("logDefault", nil)

	logDefault("logDefault 1")
	logDefault("logDefault 2", "something else")

	logStdout := dlog.New("logStdout", &dlog.Option{
		Writer: os.Stdout,
	})

	logStdout("logStdout 1", "something else", "something else")
	logStdout("logStdout 2", "something else", "something else", "something else")
}
