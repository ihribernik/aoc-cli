/*
Copyright © 2025 ivan hribernik cihribernik@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/ihribernik/aoc-cli/internal/solutions"
	_ "github.com/ihribernik/aoc-cli/internal/solutions/y2015"
	common "github.com/ihribernik/aoc-cli/pkg/common"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [year] [day]",
	Short: "Ejecuta el ejercicio para el Año dia indicado",
	Args:  cobra.ExactArgs(2),
	Run:   runner,
}

func parseArgs(args []string) (int, int, error) {
	year, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, 0, errors.New("error: year must be a string")
	}

	day, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, 0, errors.New("error: day must be a string")
	}
	return year, day, nil
}

func runner(cmd *cobra.Command, args []string) {
	year, day, err := parseArgs(args)
	if err != nil {
		log.Fatalf("Error parsing arguments: %v", err)
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
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
