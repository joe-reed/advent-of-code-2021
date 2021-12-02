package main

func puzzle1(input []int) (result int) {
	for i, measurement := range input {
		if i == 0 {
			continue
		}

		if measurement > input[i-1] {
			result += 1
		}
	}
	return
}

func puzzle2(input []int) (result int) {
	for i := range input {
		if i == len(input)-3 {
			break
		}

		if input[i+1]+input[i+2]+input[i+3] > input[i]+input[i+1]+input[i+2] {
			result += 1
		}
	}
	return
}
