package dlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	separator = ":" // or ▶
)

// terminal colors
const (
	black uint8 = iota + 30
	red
	green
	yellow
	blue
	purple
	teal
)

// rank
const (
	noRank uint8 = iota
	errorRank
	warnRank
	infoRank
	debugRank
)

var rank = errorRank

// as long as DLOG env is not empty -> show log
func init() {
	switch strings.ToUpper(os.Getenv("DLOG")) {
	case "DEBUG", "*":
		rank = debugRank
	case "INFO":
		rank = infoRank
	case "WARN":
		rank = warnRank
	case "ERROR", "ERR":
		rank = errorRank
	default:
		rank = noRank
	}
}

// Dlog .
type Dlog struct {
	name string

	prevTime   time.Time
	prevTimeMu sync.Mutex

	log    func(uint8, *Log)
	writer io.Writer
	hook   chan<- *Log

	Debug            handler
	Info, Done, Fail handler
	Warn             handler
	Error            handler
}

type handler func(...interface{})

// or should we send it to dev/null ?
func logNull(...interface{}) {}

// Option .
type Option struct {
	Writer io.Writer
	Hook   chan<- *Log
	Type   string
}

// New .
func New(name string, opt *Option) (_dlog *Dlog) {
	// blank option as default
	if opt == nil {
		opt = &Option{}
	}

	// new dlog
	_dlog = &Dlog{
		name: name,

		prevTime: time.Now(),

		writer: opt.Writer,
		hook:   opt.Hook,
	}

	// writer
	if _dlog.writer == nil {
		_dlog.writer = os.Stdout
	}

	// log
	switch opt.Type {
	case "json":
		_dlog.log = _dlog.writeJSON
	case "simple":
		_dlog.log = _dlog.writeSimple
	default:
		_dlog.log = _dlog.write
	}

	// default handler
	_dlog.Debug = logNull
	_dlog.Info, _dlog.Done, _dlog.Fail = logNull, logNull, logNull
	_dlog.Done = logNull
	_dlog.Warn = logNull
	_dlog.Error = logNull

	if rank >= debugRank {
		_dlog.Debug = _dlog.handlerFunc(teal)
	}
	if rank >= infoRank {
		_dlog.Info = _dlog.handlerFunc(blue)
		_dlog.Done = _dlog.handlerFunc(green)
		_dlog.Fail = _dlog.handlerFunc(purple)
	}
	if rank >= warnRank {
		_dlog.Warn = _dlog.handlerFunc(yellow)
	}
	if rank >= errorRank {
		_dlog.Error = _dlog.handlerFunc(red)
	}

	return
}

func (d *Dlog) handlerFunc(color uint8) handler {
	return func(arg ...interface{}) {
		// time
		d.prevTimeMu.Lock()
		now, delta := getDelta(d.prevTime)
		d.prevTime = now
		d.prevTimeMu.Unlock()

		_log := NewLog(d.name, now, delta, arg...)

		// write to writer
		d.log(color, _log)

		// send to hook
		if d.hook != nil {
			d.hook <- _log
		}
	}
}

func (d *Dlog) write(color uint8, _log *Log) {
	timestamp := _log.Timestamp.Format("15:04:05.000")
	prefix := fmt.Sprintf("\033[%vm%v %s #%v %v\033[0m", color, timestamp, _log.Name, _log.Caller, separator)
	delta := fmt.Sprintf("\033[%vm+%s\033[0m", black, humanizeNano(_log.Delta))

	arg := append([]interface{}{prefix}, _log.Data...)
	arg = append(arg, delta)

	// skip err
	fmt.Fprintln(d.writer, arg...)
}

// no color, no time, no delta
func (d *Dlog) writeSimple(color uint8, _log *Log) {
	prefix := fmt.Sprintf("%s #%v %v", _log.Name, _log.Caller, separator)

	arg := append([]interface{}{prefix}, _log.Data...)

	// skip err
	fmt.Fprintln(d.writer, arg...)
}

func (d *Dlog) writeJSON(color uint8, _log *Log) {
	jsonStr, err := json.Marshal(_log)
	// skip err
	if err != nil {
		return
	}

	// skip err
	fmt.Fprintln(d.writer, string(jsonStr))
	return
}
