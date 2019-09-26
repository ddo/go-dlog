package caller

import (
	"testing"
)

var caseTrim = [][]string{
	{"github.com/ddo/request.New", "New"},
	{"github.com/ddo/request.(*Client).Request", "Request"},
	{"main.main", "main"},
}

func TestTrim(t *testing.T) {
	for i := 0; i < len(caseTrim); i++ {
		if trim(caseTrim[i][0]) != caseTrim[i][1] {
			t.Error(caseTrim[i][0] + " should be " + caseTrim[i][1])
		}
	}
}

// funcName is goexit
func TestGet(t *testing.T) {
	if Get() != "goexit" {
		t.Error()
	}
}
