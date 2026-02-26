package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/ihribernik/aoc-cli/internal/container"
	runusecase "github.com/ihribernik/aoc-cli/internal/run"
	"github.com/spf13/cobra"
)

var appContainer = container.New()

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:          "run --year <year> --day <day>",
	Short:        "Execute the solution for a specific Advent of Code puzzle",
	Args:         cobra.NoArgs,
	RunE:         runE,
	SilenceUsage: true,
}

func runE(cmd *cobra.Command, args []string) error {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		return fmt.Errorf("invalid --year value: %w", err)
	}

	day, err := cmd.Flags().GetInt("day")
	if err != nil {
		return fmt.Errorf("invalid --day value: %w", err)
	}

	if day < 1 || day > 25 {
		return fmt.Errorf("invalid --day %d: expected a value between 1 and 25", day)
	}

	result, err := appContainer.Runner.Execute(year, day)
	if err != nil {
		return mapRunError(year, day, err)
	}

	fmt.Println("Solution part 1:", result.Part1)
	fmt.Println("Solution part 2:", result.Part2)
	return nil
}

func mapRunError(year int, day int, err error) error {
	var registerErr *runusecase.ErrRegisterYear
	if errors.As(err, &registerErr) {
		return fmt.Errorf("year %d is not available: %w", registerErr.Year, registerErr.Err)
	}

	var solverErr *runusecase.ErrSolverNotFound
	if errors.As(err, &solverErr) {
		return fmt.Errorf("no solver registered for year %d day %02d", solverErr.Year, solverErr.Day)
	}

	var inputErr *runusecase.ErrGetInput
	if errors.As(err, &inputErr) {
		if errors.Is(inputErr.Err, os.ErrNotExist) || os.IsNotExist(inputErr.Err) {
			return fmt.Errorf("input file not found for year %d day %02d", inputErr.Year, inputErr.Day)
		}
		return fmt.Errorf("cannot load input for year %d day %02d: %w", inputErr.Year, inputErr.Day, inputErr.Err)
	}

	var solveErr *runusecase.ErrSolvePart
	if errors.As(err, &solveErr) {
		return fmt.Errorf("failed while solving part %d: %w", solveErr.Part, solveErr.Err)
	}

	return fmt.Errorf("run failed for year %d day %02d: %w", year, day, err)
}

func init() {
	runCmd.Flags().IntP("year", "y", 0, "Year to execute (e.g. 2015)")
	runCmd.Flags().IntP("day", "d", 0, "Day to execute (e.g. 6)")
	_ = runCmd.MarkFlagRequired("year")
	_ = runCmd.MarkFlagRequired("day")
	rootCmd.AddCommand(runCmd)
}
