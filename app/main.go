package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ddo/go-dlog"
)

func main() {
	// create dlog
	opt := dlog.Option{
		Writer: os.Stdout,
	}

	// check whether output is terminal
	// if not, log type is json
	if !dlog.IsTTY(os.Stdout) {
		opt.Type = "json"
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		line, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		str := string(line)
		fmt.Println(str)
		fmt.Println(fmt.Sprintf("%q", str))

		fmt.Println(isPrefix)
		fmt.Println("---------------")
	}

	fmt.Println("DONE app")
}
