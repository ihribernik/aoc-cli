package common

import "slices"

func ContainsIn(values []string, valuesToCompare []string) bool {
	for _, valueToCompare := range valuesToCompare {
		if slices.Contains(values, valueToCompare) {
			return true
		}
	}
	return false
}
