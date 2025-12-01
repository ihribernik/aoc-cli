/*
Copyright © 2025 ivan hribernik cihribernik@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ihribernik/aoc-cli/internal/solutions"
	_ "github.com/ihribernik/aoc-cli/internal/solutions/y2015"
	common "github.com/ihribernik/aoc-cli/pkg/common"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run --year <year> --day <day>",
	Short: "Ejecuta el ejercicio para el Año dia indicado",
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

	solver, ok := solutions.GetSolver(year, day)
	if !ok {
		log.Fatalf("No se encontró solución para el año %d, día %d", year, day)
	}

	input, err := common.GetInput(year, day)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	result1, err := solver.SolvePart1(input)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Parte 1:", result1)

	result2, err := solver.SolvePart2(input)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Parte 2:", result2)
}

func init() {
	runCmd.Flags().IntP("year", "y", 0, "Año del ejercicio (ej: 2015)")
	runCmd.Flags().IntP("day", "d", 0, "Día del ejercicio (ej: 6)")
	_ = runCmd.MarkFlagRequired("year")
	_ = runCmd.MarkFlagRequired("day")
	rootCmd.AddCommand(runCmd)
}
