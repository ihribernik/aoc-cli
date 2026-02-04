package y2015

type Day01 struct{}

func (d Day01) SolvePart1(input []string) (int, error) {
	result := 0
	for _, v := range input {
		switch v {
		case "(":
			result++
		case ")":
			result--
		default:
			continue
		}
	}

	return result, nil
}

func (d Day01) SolvePart2(input []string) (int, error) {
	positionChar := 0
	floor := 0

	for index, v := range input {
		if floor == -1 {
			positionChar = index
			break
		}

		switch v {
		case "(":
			floor++
		case ")":
			floor--
		default:
			continue
		}
	}

	return positionChar, nil
}
