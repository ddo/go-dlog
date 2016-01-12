package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ddo/go-dlog"
	"gopkg.in/ddo/request.v1"
)

var hook chan *dlog.Log
var client *request.Client

func init() {
	hook = make(chan *dlog.Log)
	client = request.New()
}

func main() {
	logStdoutHook := dlog.New("logStdoutHook", &dlog.Option{
		Hook: hook,
	})

	go (func() {
		for {
			log := <-hook
			go postToSlack(log)
		}
	})()

	logStdoutHook("logStdoutHook 1")
	logStdoutHook("logStdoutHook 2")

	// wait for hook output before main terminate
	time.Sleep(time.Second * 10)
}

func postToSlack(log *dlog.Log) {
	// u should handle your log.Data before send it as text
	// in this case i just use the 1st data text
	text := log.Data[0]

	res, err := client.Request(&request.Option{
		Url:    os.Getenv("SLACK_DEV_HOOK"),
		Method: "POST",
		Json: map[string]interface{}{
			"username":   "dLog",
			"icon_emoji": ":robot_face:",
			"text":       text,
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()
}
