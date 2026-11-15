package cmd

import (
	"fmt"

	runusecase "github.com/ihribernik/aoc-cli/internal/run"
	"github.com/spf13/cobra"
)

func parseRunOptions(cmd *cobra.Command) (int, int, error) {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		return 0, 0, fmt.Errorf("invalid --year value: %w", err)
	}

	day, err := cmd.Flags().GetInt("day")
	if err != nil {
		return 0, 0, fmt.Errorf("invalid --day value: %w", err)
	}

	if day < 1 || day > 25 {
		return 0, 0, fmt.Errorf("invalid --day %d: expected a value between 1 and 25", day)
	}
	return year, day, nil
}

func newRunCmd(runner runusecase.Runner) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run --year <year> --day <day>",
		Short: "Execute the solution for a specific Advent of Code puzzle",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			year, day, err := parseRunOptions(cmd)
			if err != nil {
				return err
			}

			if runner == nil {
				return newCLIError(
					"application runner is not configured",
					runusecase.ErrRunnerNotConfigured,
				)
			}

			result, err := runner.Execute(year, day)
			if err != nil {
				return mapRunError(year, day, err)
			}

			if _, err := fmt.Fprintln(
				cmd.OutOrStdout(), "Solution part 1:", result.Part1,
			); err != nil {
				return err
			}

			_, err = fmt.Fprintln(
				cmd.OutOrStdout(), "Solution part 2:", result.Part2,
			)
			return err
		},
		SilenceUsage: true,
	}
	cmd.Flags().IntP("year", "y", 0, "Year to execute (e.g. 2015)")
	cmd.Flags().IntP("day", "d", 0, "Day to execute (e.g. 6)")
	_ = cmd.MarkFlagRequired("year")
	_ = cmd.MarkFlagRequired("day")
	return cmd
}
