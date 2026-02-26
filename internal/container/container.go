package container

import (
	"github.com/ihribernik/aoc-cli/internal/inputs"
	"github.com/ihribernik/aoc-cli/internal/registry"
	"github.com/ihribernik/aoc-cli/internal/run"
	"github.com/ihribernik/aoc-cli/internal/solutions"
)

type Container struct {
	Runner *run.Runner
}

func New() *Container {
	return &Container{
		Runner: run.NewRunner(
			registry.NewRegistry,
			solutions.RegisterYear,
			inputs.GetInput,
		),
	}
}
