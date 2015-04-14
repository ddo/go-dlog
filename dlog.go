package dlog

import (
	"log"
	"os"
	"time"
)

// round robin color
var i = 0

// terminal colors
var colors = []string{
	"31",
	"32",
	"33",
	"34",
	"35",
	"36",
}

var enabled = false

// as long as DLOG env is not empty -> show log
func init() {
	if os.Getenv("DLOG") != "" {
		enabled = true
	}
}

func New(name string) func(...interface{}) {
	if !enabled {
		return func(...interface{}) {}
	}

	color := colors[i%len(colors)]
	i++

	logger := log.New(os.Stdout, "", 0)

	return func(arg ...interface{}) {
		timestamp := time.Now().Format("15:04:05.000")
		fmt := "\033[" + color + "m" + timestamp + " " + name + "\t▶\033[0m"

		arg = append([]interface{}{fmt}, arg...)
		logger.Println(arg...)
	}
}
