package main

import (
	"strconv"
	. "utils"
)

func puzzle1(input []string) int {
	gammaRate := ""
	epsilonRate := ""

	for i := 0; i < len(input[0]); i++ {
		mostCommon, leastCommon := getMostAndLeastCommonBits(input, i)

		gammaRate += mostCommon
		epsilonRate += leastCommon
	}

	return binaryToDecimal(gammaRate) * binaryToDecimal(epsilonRate)
}

func puzzle2(input []string) int {
	oxygenCandidates := make([]string, len(input))
	co2Candidates := make([]string, len(input))

	copy(oxygenCandidates, input)
	copy(co2Candidates, input)

	for i := 0; i < len(input[0]); i++ {
		if len(oxygenCandidates) == 1 {
			break
		}

		newOxygenCandidates := []string{}

		mostCommon, _ := getMostAndLeastCommonBits(oxygenCandidates, i)

		for _, line := range oxygenCandidates {
			if string(line[i]) == mostCommon {
				newOxygenCandidates = append(newOxygenCandidates, line)
			}
		}

		oxygenCandidates = newOxygenCandidates
	}

	for i := 0; i < len(input[0]); i++ {
		if len(co2Candidates) == 1 {
			break
		}

		newCo2Candidates := []string{}

		_, leastCommon := getMostAndLeastCommonBits(co2Candidates, i)

		for _, line := range co2Candidates {
			if string(line[i]) == leastCommon {
				newCo2Candidates = append(newCo2Candidates, line)
			}
		}

		co2Candidates = newCo2Candidates
	}

	return binaryToDecimal(oxygenCandidates[0]) * binaryToDecimal(co2Candidates[0])
}

func getMostAndLeastCommonBits(input []string, position int) (string, string) {
	zeroCount := 0
	oneCount := 0

	for _, line := range input {
		if string(line[position]) == "0" {
			zeroCount++
		} else {
			oneCount++
		}
	}

	if zeroCount > oneCount {
		return "0", "1"
	}

	return "1", "0"
}

func binaryToDecimal(binary string) int {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	Check(err)
	return int(decimal)
}
