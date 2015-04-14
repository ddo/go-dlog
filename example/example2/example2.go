package example2

import (
	"github.com/ddo/go-dlog"
)

func Test(name string) {
	log := dlog.New(name)
	log("some log")
}
