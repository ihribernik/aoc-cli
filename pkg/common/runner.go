package common

import (
	"fmt"
	"os"
	"strings"

	"path/filepath"
)

func GetInput(year int, day int) ([]string, error) {
	filepath := filepath.Join("input", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d.txt", day))
	data, err := os.ReadFile(filepath)
	if err != nil {
		return []string{}, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	lineas := strings.Split(content, "\n")

	return lineas, nil
}
