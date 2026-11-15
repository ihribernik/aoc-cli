package cmd

import (
	"github.com/ihribernik/aoc-cli/internal/container"
	runusecase "github.com/ihribernik/aoc-cli/internal/run"
	"github.com/spf13/cobra"
)

func newRootCmd(runner runusecase.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "aoc-cli",
		Short:         "Advent of Code CLI runner",
		Long:          `Advent of Code CLI is a command-line tool to run Advent of Code solutions.`,
		SilenceErrors: true,
	}
	cmd.AddCommand(newRunCmd(runner))

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	configuredContainer, err := container.New()
	if err != nil {
		return mapStartupError(err)
	}

	cmd := newRootCmd(configuredContainer.GetRunner())
	return cmd.Execute()
}
