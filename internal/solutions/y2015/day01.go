package y2015

type Day01 struct{}

func (d Day01) SolvePart1(input []string) (int, error) {
	floor := 0

	for _, line := range input {
		for _, ch := range line {
			switch ch {
			case '(':
				floor++
			case ')':
				floor--
			}
		}
	}

	return floor, nil
}

func (d Day01) SolvePart2(input []string) (int, error) {
	position := 0
	floor := 0

	for _, line := range input {
		for _, ch := range line {
			position++
			switch ch {
			case '(':
				floor++
			case ')':
				floor--
			}

			if floor == -1 {
				return position, nil
			}
		}
	}

	return 0, nil
}
