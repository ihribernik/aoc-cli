package y2015

import (
	"slices"
	"strings"

	"github.com/ihribernik/aoc-cli/internal/directions"
)

type Day03 struct{}

func (d Day03) parseLine(input []string) []string {
	inputVars := strings.Split(input[0], "")
	return inputVars
}

func (d Day03) SolvePart1(input []string) (int, error) {
	position := directions.NewDirection(0, 0)

	visited := map[directions.Direction]struct{}{
		position: {},
	}

	characters := d.parseLine(input)

	for _, direction := range characters {
		currentDirection := directions.DIRECTIONS[direction]
		position = currentDirection.NewPositionWith(position)
		visited[position] = struct{}{}
	}
	return len(visited), nil
}

func (d Day03) SolvePart2(input []string) (int, error) {
	position := directions.NewDirection(0, 0)

	positions := []directions.Direction{
		position,
		position,
	}

	visited := map[directions.Direction]struct{}{
		position: {},
	}

	inputDirections := d.parseLine(input)

	for chunk := range slices.Chunk(inputDirections, 2) {
		for i, direction := range chunk {
			currentDirection := directions.DIRECTIONS[direction]
			positions[i] = currentDirection.NewPositionWith(positions[i])
			visited[positions[i]] = struct{}{}
		}
	}
	return len(visited), nil
}
