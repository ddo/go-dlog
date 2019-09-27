package dlog

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// IsTTY returns whether file is terminal or not
func IsTTY(file *os.File) bool {
	return terminal.IsTerminal(int(file.Fd()))
}
