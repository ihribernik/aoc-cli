/*
Copyright © 2025 ivan hribernik cihribernik@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc-cli",
	Short: "Advent of code cli runner",
	Long:  `Advent of code cli es una herramienta de línea de comandos para resolver los ejercicios de Advent of Code.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
