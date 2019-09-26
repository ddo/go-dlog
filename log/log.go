package log

import (
	"time"

	"github.com/ddo/go-dlog/caller"
)

// Log is a single log object
type Log struct {
	Rank string `json:"rank"`

	Name   string `json:"name"`
	Caller string `json:"caller"` // function name

	Timestamp time.Time `json:"timestamp"`

	Data []interface{} `json:"data"`
}

// New returns new Log object
func New(rank, name string, now time.Time, arg ...interface{}) *Log {
	return &Log{
		Rank: rank,

		Name:   name,
		Caller: caller.Get(),

		Timestamp: now,

		Data: arg,
	}
}
