package dlog

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/ddo/go-dlog/log"
)

const (
	separator = ":" // or ▶
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

var colors = map[string]uint8{
	"DEBUG": teal,
	"INFO":  blue,
	"DONE":  green,
	"FAIL":  purple,
	"WARN":  yellow,
	"ERROR": red,
}

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

// Dlog is the logger object
type Dlog struct {
	name string

	log    func(*log.Log)
	writer io.Writer
	hook   chan<- *log.Log

	Debug            handler
	Info, Done, Fail handler
	Warn             handler
	Error            handler
}

type handler func(...interface{})

// or should we send it to dev/null ?
func logNull(...interface{}) {}

// Option is the option for #New
type Option struct {
	Writer io.Writer
	Hook   chan<- *log.Log
	Type   string
}

// New returns the Dlog object
func New(name string, opt *Option) (_dlog *Dlog) {
	// blank option as default
	if opt == nil {
		opt = &Option{}
	}

	// writer
	// default is stdout
	writer := opt.Writer
	if opt.Writer == nil {
		writer = os.Stdout
	}

	// type
	// default is json
	// but if writer is stdout and tty then pretty
	_type := opt.Type
	if _type == "" && writer == os.Stdout && IsTTY(os.Stdout) {
		_type = "pretty"
	}
	if _type == "" {
		_type = "json"
	}

	// new dlog
	_dlog = &Dlog{
		name: name,

		writer: writer,
		hook:   opt.Hook,
	}

	// type
	switch _type {
	case "json":
		_dlog.log = _dlog.WriteJSON

	case "simple":
		_dlog.log = _dlog.WriteSimple

	default:
		_dlog.log = _dlog.Write
	}

	// default handler
	_dlog.Debug = logNull
	_dlog.Info, _dlog.Done, _dlog.Fail = logNull, logNull, logNull
	_dlog.Done = logNull
	_dlog.Warn = logNull
	_dlog.Error = logNull

	if rank >= debugRank {
		_dlog.Debug = _dlog.handlerFunc("DEBUG")
	}

	if rank >= infoRank {
		_dlog.Info = _dlog.handlerFunc("INFO")
		_dlog.Done = _dlog.handlerFunc("DONE")
		_dlog.Fail = _dlog.handlerFunc("FAIL")
	}

	if rank >= warnRank {
		_dlog.Warn = _dlog.handlerFunc("WARN")
	}

	if rank >= errorRank {
		_dlog.Error = _dlog.handlerFunc("ERROR")
	}

	return
}

func (d *Dlog) handlerFunc(rank string) handler {
	return func(arg ...interface{}) {
		_log := log.New(rank, d.name, time.Now(), arg...)

		// write to writer
		d.log(_log)

		// send to hook
		if d.hook != nil {
			d.hook <- _log
		}
	}
}

// Write writes pretty log with colors
func (d *Dlog) Write(_log *log.Log) {
	timestamp := _log.Timestamp.Format("15:04:05.000")
	prefix := fmt.Sprintf("\033[%vm%v %s #%v %v\033[0m", colors[_log.Rank], timestamp, _log.Name, _log.Caller, separator)

	arg := append([]interface{}{prefix}, _log.Data...)

	// skip err
	fmt.Fprintln(d.writer, arg...)
}

// WriteSimple writes simple log with no color, no time
func (d *Dlog) WriteSimple(_log *log.Log) {
	prefix := fmt.Sprintf("%-5s %s #%v %v", _log.Rank, _log.Name, _log.Caller, separator)

	arg := append([]interface{}{prefix}, _log.Data...)

	// skip err
	fmt.Fprintln(d.writer, arg...)
}

// WriteJSON writes log as json
func (d *Dlog) WriteJSON(_log *log.Log) {
	jsonStr, err := json.Marshal(_log)
	// skip err
	if err != nil {
		return
	}

	// skip err
	fmt.Fprintln(d.writer, string(jsonStr))
	return
}
