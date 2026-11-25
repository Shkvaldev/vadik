package tools

import (
	"bufio"
	"fmt"
	"os"
)

// StdAskCallback is a function for asking permissions from stdin
func StdAskCallback(question string) bool {
	fmt.Printf("[ALLOW?] %s (type `y` to accept):\n", question)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return false
	}

	if res := scanner.Text(); res != "y" {
		return false
	}

	return true
}
