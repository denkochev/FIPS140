package fips140

func Runs(bits_set []uint64) (bool, map[int]int) {
	/*
	   table from this specs [pg.45]
	   https://csrc.nist.gov/files/pubs/fips/140-1/upd1/final/docs/fips1401.pdf
	*/
	specsTable := map[int][2]int{
		1: [2]int{2267, 2733},
		2: [2]int{1079, 1421},
		3: [2]int{502, 748},
		4: [2]int{223, 402},
		5: [2]int{90, 223},
		6: [2]int{90, 223},
	}

	seriesOnes, seriesZeros := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}, map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}

	curStreak := 0
	curBit := -1

	block := 64

	for i := 0; i < len(bits_set); i++ {
		curBlock := bits_set[i]

		if i == len(bits_set)-1 {
			block = 32
		}

		for j := block - 1; j >= 0; j-- {
			highest_bit := int((curBlock >> j) & 1)
			if curBit == -1 {
				curBit = highest_bit
				curStreak = 1
			} else if curBit == highest_bit {
				curStreak += 1
			} else {
				if curStreak >= 6 {
					if curBit == 1 {
						seriesOnes[6] += 1
					} else {
						seriesZeros[6] += 1
					}
				} else {
					if curBit == 1 {
						seriesOnes[curStreak] += 1
					} else {
						seriesZeros[curStreak] += 1
					}
				}

				curBit = highest_bit
				curStreak = 1
			}
		}
	}

	// add stash to the table
	if curStreak >= 6 {
		if curBit == 1 {
			seriesOnes[6] += 1
		} else {
			seriesZeros[6] += 1
		}
	} else {
		if curBit == 1 {
			seriesOnes[curStreak] += 1
		} else {
			seriesZeros[curStreak] += 1
		}
	}

	result := true

	for i := 1; i <= 6; i++ {
		curSequenceOnes, curSequenceZeros := seriesOnes[i], seriesZeros[i]
		specsMin, specsMax := specsTable[i][0], specsTable[i][1]

		if specsMin > curSequenceOnes || specsMax < curSequenceOnes || specsMin > curSequenceZeros || specsMax < curSequenceZeros {
			result = false
		}
	}

	return result, seriesOnes
}
