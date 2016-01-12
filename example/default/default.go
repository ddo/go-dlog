package main

import (
	"os"

	"github.com/ddo/go-dlog"
)

// if nil writer -> os.Stdout
func main() {
	logEmail := dlog.New("email", nil)
	logEmail("send to hi@gmail.com", "support@gmail.com")

	// same as default = writer = os.Stdout
	logHttp := dlog.New("http", &dlog.Option{
		Writer: os.Stdout,
	})

	logHttp("GET", "/orders", "200", "OK")
	logHttp("POST", "/order", "201", "new order")
	logEmail("send to customer@gmail.com")
}
