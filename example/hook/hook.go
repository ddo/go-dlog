package main

import (
	"fmt"
	"time"

	"github.com/ddo/go-dlog"
)

func main() {
	hook := make(chan *dlog.Log)

	logStdoutHook := dlog.New("logStdoutHook", &dlog.Option{
		Hook: hook,
	})

	go (func() {
		for {
			log := <-hook
			fmt.Println("logStdoutHook: ", log)
		}
	})()

	logStdoutHook("logStdoutHook 1")
	logStdoutHook("logStdoutHook 2")

	// wait for hook output before main terminate
	time.Sleep(time.Second)
}
