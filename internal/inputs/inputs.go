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
	if strings.HasSuffix(content, "\n") {
		content = strings.TrimSuffix(content, "\n")
	}
	if content == "" {
		return []string{}, nil
	}

	lines := strings.Split(content, "\n")
	return lines, nil
}
