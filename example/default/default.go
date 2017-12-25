package main

import (
	"os"

	"github.com/ddo/go-dlog"
)

// if nil writer -> os.Stdout
func main() {
	logEmail := dlog.New("email", nil)
	logEmail.Info("send to hi@gmail.com", "support@gmail.com")

	// same as default = writer = os.Stdout
	LogHTTP := dlog.New("http", &dlog.Option{
		Writer: os.Stdout,
	})

	LogHTTP.Debug("GET", "/orders", "200", "OK")
	LogHTTP.Debug("POST", "/order", "201", "new order")
	logEmail.Info("send to customer@gmail.com")
}
