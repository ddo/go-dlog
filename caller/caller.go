package caller

import (
	"runtime"
	"strings"
)

// Get returns the runtime caller name
func Get() (caller string) {
	caller = ""

	pc := make([]uintptr, 1)

	if runtime.Callers(4, pc) == 1 {
		f := runtime.FuncForPC(pc[0])
		caller = trim(f.Name())
	}

	return
}

// "github.com/ddo/request.(*Client).Request"
// -> (*Client).Request
// or
// "github.com/ddo/request.New"
// -> New
func trim(funcName string) string {
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
