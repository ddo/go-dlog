package dlog

import (
	"testing"
)

var caseTrimCaller = [][]string{
	{"github.com/ddo/request.New", "New"},
	{"github.com/ddo/request.(*Client).Request", "Request"},
	{"main.main", "main"},
}

func TestTrimCaller(t *testing.T) {
	for i := 0; i < len(caseTrimCaller); i++ {
		if trimCaller(caseTrimCaller[i][0]) != caseTrimCaller[i][1] {
			t.Error(caseTrimCaller[i][0] + " should be " + caseTrimCaller[i][1])
		}
	}
}

// funcName is goexit
func TestGetCaller(t *testing.T) {
	if getCaller() != "goexit" {
		t.Error()
	}
}
