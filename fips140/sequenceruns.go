package fips140

import "fmt"

func SequencesRuns(bits_set []uint64) bool {
	series := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	maxStreak, curStreak := 0, 0

	var curValue int = -1

	block := 64

	for i := 0; i < len(bits_set); i++ {
		curBlock := bits_set[i]

		if i == len(bits_set)-1 {
			block = 32
		}

		for j := block - 1; j >= 0; j-- {
			highest_bit := int((curBlock >> j) & 1)
			if curValue == -1 {
				curValue = highest_bit
				curStreak += 1
				if curStreak > maxStreak {
					maxStreak = curStreak
				}
			} else {
				if highest_bit == curValue {
					curStreak += 1
					if curStreak > maxStreak {
						maxStreak = curStreak
					}
				} else {
					if curStreak >= 6 {
						series[6] += 1
					} else {
						series[curStreak] += 1
					}

					curValue = highest_bit
					curStreak = 0
				}
			}

		}
	}
	fmt.Println(series)
	return true
}
