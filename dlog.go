package dlog

import (
	"fmt"
	"os"
	"strconv"
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

	//color
	color := colors[i%len(colors)]
	i++

	//save for delta
	prevTime := time.Now()

	return func(arg ...interface{}) {
		now := time.Now()

		delta := now.Sub(prevTime).Nanoseconds()
		prevTime = now

		timestamp := now.Format("15:04:05.000")

		prefix := fmt.Sprintf("\033[%vm%v %-6s %-10s ▶\033[0m", color, timestamp, humanizeNano(delta), name)

		arg = append([]interface{}{prefix}, arg...)
		fmt.Println(arg...)
	}
}

func humanizeNano(n int64) string {
	var suffix string

	switch {
	case n > 1e9:
		n /= 1e9
		suffix = "s"
	case n > 1e6:
		n /= 1e6
		suffix = "ms"
	case n > 1e3:
		n /= 1e3
		suffix = "us"
	default:
		suffix = "ns"
	}

	return strconv.Itoa(int(n)) + suffix
}
