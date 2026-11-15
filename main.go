/*
Copyright © 2025 Ivan Hribernik <cihribernik@gmail.com>
*/
package main

import (
	"fmt"
	"os"

	"github.com/ihribernik/aoc-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
