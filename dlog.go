package dlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	// SEPARATOR = "▶"
	SEPARATOR = ":"
)

// round robin color
var i = 0

// terminal colors
var colors = []uint8{
	31,
	32,
	33,
	34,
	35,
	36,
}

var enabled = false

// as long as DLOG env is not empty -> show log
func init() {
	if os.Getenv("DLOG") != "" {
		enabled = true
	}
}

// Caller: function name
type Log struct {
	Name   string `json:"name"`
	Caller string `json:"caller"`

	Timestamp time.Time     `json:"timestamp"`
	Delta     time.Duration `json:"delta"`

	Data []interface{} `json:"data"`
}

func New(name string, writer io.Writer) func(...interface{}) {
	// disable dlog
	// or should we send it to dev/null ?
	if !enabled {
		return func(...interface{}) {}
	}

	// color
	color := colors[i%len(colors)]
	i++

	// save for delta
	prevTime := time.Now()

	return func(arg ...interface{}) {
		now := time.Now()

		delta := now.Sub(prevTime)
		prevTime = now

		log := &Log{
			Name:   name,
			Caller: getCaller(),

			Timestamp: now,
			Delta:     delta,

			Data: arg,
		}

		write(writer, log, color)
	}
}

func write(writer io.Writer, log *Log, color uint8) {
	if writer != nil && writer != os.Stdout {
		jsonStr, err := json.Marshal(log)

		// skip err
		if err != nil {
			return
		}

		fmt.Fprintln(writer, string(jsonStr))
		return
	}

	writer = os.Stdout

	timestamp := log.Timestamp.Format("15:04:05.000")

	prefix := fmt.Sprintf("\033[%vm%v %-6s %-10s #%v %v\033[0m", color, timestamp, humanizeNano(log.Delta), log.Name, log.Caller, SEPARATOR)

	arg := append([]interface{}{prefix}, log.Data...)

	fmt.Fprintln(writer, arg...)
}

func getCaller() (caller string) {
	caller = "unknown"

	pc := make([]uintptr, 1)

	if runtime.Callers(3, pc) == 1 {
		f := runtime.FuncForPC(pc[0])
		callers := strings.Split(f.Name(), ".")

		if len(callers) > 0 {
			caller = callers[len(callers)-1]
		}
	}

	return
}

func humanizeNano(n time.Duration) string {
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
