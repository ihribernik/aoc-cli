package inputs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetInput(year int, day int) ([]string, error) {
	path := filepath.Join("input", fmt.Sprintf("%d", year), fmt.Sprintf("day%02d.txt", day))
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	content := strings.NewReplacer("\r\n", "\n", "\r", "\n").Replace(string(data))
	content = strings.TrimRight(content, "\n")
	if content == "" {
		return []string{}, ErrEmptyInput
	}

	return strings.Split(content, "\n"), nil
}
