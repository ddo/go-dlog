package dlog

import (
	"testing"
)

var caseTrimCaller = [][]string{
	[]string{"github.com/ddo/request.New", "New"},
	[]string{"github.com/ddo/request.(*Client).Request", "(*Client).Request"},
	[]string{"main.main", "main"},
}

func TestTrimCaller(t *testing.T) {
	for i := 0; i < len(caseTrimCaller); i++ {
		if trimCaller(caseTrimCaller[i][0]) != caseTrimCaller[i][1] {
			t.Error(caseTrimCaller[i][0] + " should be " + caseTrimCaller[i][1])
		}
	}
}

// funcName is testing.tRunner
func TestGetCaller(t *testing.T) {
	if getCaller() != "tRunner" {
		t.Error()
	}
}
