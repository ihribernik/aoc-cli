package cmd

import (
	"fmt"
	"log"

	"github.com/ihribernik/aoc-cli/internal/inputs"
	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/solutions"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run --year <year> --day <day>",
	Short: "Execute the solution for a specific puzzle in the year day requested",
	Args:  cobra.NoArgs,
	Run:   runner,
}

func runner(cmd *cobra.Command, args []string) {
	year, err := cmd.Flags().GetInt("year")
	if err != nil {
		log.Fatalf("Error parsing year: %v", err)
	}

	day, err := cmd.Flags().GetInt("day")
	if err != nil {
		log.Fatalf("Error parsing day: %v", err)
	}

	reg := registry.NewRegistry()

	err = solutions.RegisterYear(reg, year)
	if err != nil {
		log.Fatalf("failed to register the year %d, err: %d", year, err)
	}

	solver, ok := reg.GetSolver(year, day)

	if !ok {
		log.Fatalf("cannot find a solution for the year %d, day %d", year, day)
	}

	input, err := inputs.GetInput(year, day)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result1, err := solver.SolvePart1(input)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Solution part 1:", result1)

	result2, err := solver.SolvePart2(input)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("solution part 2:", result2)
}

func init() {
	runCmd.Flags().IntP("year", "y", 0, "Year to execute (ej: 2015)")
	runCmd.Flags().IntP("day", "d", 0, "Day to execute (ej: 6)")
	_ = runCmd.MarkFlagRequired("year")
	_ = runCmd.MarkFlagRequired("day")
	rootCmd.AddCommand(runCmd)
}
