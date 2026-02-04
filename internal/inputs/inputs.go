package inputs

import (
	"fmt"
	"os"
	"strings"
)

func GetInput(year int, day int) ([]string, error) {
	filepath := fmt.Sprintf("input/%d/day%02d.txt", year, day)
	data, err := os.ReadFile(filepath)
	if err != nil {
		return []string{}, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")

	lineas := strings.Split(content, "\n")

	return lineas, nil
}
