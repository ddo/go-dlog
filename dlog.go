package dlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
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

type Option struct {
	Writer io.Writer
	Hook   chan<- *Log
}

// Caller: function name
type Log struct {
	Name   string `json:"name"`
	Caller string `json:"caller"`

	Timestamp time.Time     `json:"timestamp"`
	Delta     time.Duration `json:"delta"`

	Data []interface{} `json:"data"`
}

func New(name string, opt *Option) func(...interface{}) {
	// disable dlog
	// or should we send it to dev/null ?
	if !enabled {
		return func(...interface{}) {}
	}

	// blank option as default
	if opt == nil {
		opt = &Option{}
	}

	// color
	color := colors[i%len(colors)]
	i++

	// sync
	mutex := sync.Mutex{}

	// save for delta
	prevTime := time.Now()

	return func(arg ...interface{}) {
		now := time.Now()

		mutex.Lock()

		delta := now.Sub(prevTime)
		prevTime = now

		mutex.Unlock()

		log := &Log{
			Name:   name,
			Caller: getCaller(),

			Timestamp: now,
			Delta:     delta,

			Data: arg,
		}

		// writer to writer
		write(opt.Writer, log, color)

		// send to hook
		if opt.Hook != nil {
			opt.Hook <- log
		}
	}
}

func write(writer io.Writer, log *Log, color uint8) {
	if writer != nil && writer != os.Stdout {
		jsonStr, err := json.Marshal(log)

		// skip err
		if err != nil {
			return
		}

		// skip err
		fmt.Fprintln(writer, string(jsonStr))
		return
	}

	writer = os.Stdout

	timestamp := log.Timestamp.Format("15:04:05.000")

	prefix := fmt.Sprintf("\033[%vm%v %-6s %-10s #%v %v\033[0m", color, timestamp, humanizeNano(log.Delta), log.Name, log.Caller, SEPARATOR)

	arg := append([]interface{}{prefix}, log.Data...)

	// skip err
	fmt.Fprintln(writer, arg...)
}

func getCaller() (caller string) {
	caller = ""

	pc := make([]uintptr, 1)

	if runtime.Callers(3, pc) == 1 {
		f := runtime.FuncForPC(pc[0])
		caller = trimCaller(f.Name())
	}

	return
}

// "github.com/ddo/request.(*Client).Request"
// -> (*Client).Request
// or
// "github.com/ddo/request.New"
// -> New
func trimCaller(funcName string) string {
	// ex:
	// funcName = "github.com/ddo/request.(*Client).Request"
	// arrDir = [github.com ddo request.(*Client).Request]
	// lastDir = request.(*Client).Request
	// arrCaller = [request. (*Client).Request]

	arrDir := strings.Split(funcName, "/")
	lastDir := arrDir[len(arrDir)-1]
	arrCaller := strings.Split(lastDir, ".")

	if len(arrCaller) < 2 {
		return ""
	}

	return arrCaller[len(arrCaller)-1]
}

// copty from TJ
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
