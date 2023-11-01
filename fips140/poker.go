package fips140

import (
	"math"
)

func Poker(bits_set []uint64) bool {
	fourBitsSeries := map[uint64]int{}

	block := 64

	for i := 0; i < len(bits_set); i++ {
		curBlock := bits_set[i]

		if i == len(bits_set)-1 {
			block = 32
		}

		for j := 0; j < block; j += 4 {
			fourBits := curBlock & 0x0F
			fourBitsSeries[fourBits] += 1
			curBlock = curBlock >> 4
		}
	}

	// CALCULATE SUM FROM POKER FORMULA
	var sum float64 = 0
	var i uint64
	for i = 0; i < 16; i++ {
		sum += math.Pow(float64(fourBitsSeries[i]), 2)
	}
	// FORMULA [pg.43] https://csrc.nist.gov/files/pubs/fips/140-1/upd1/final/docs/fips1401.pdf
	poker := (16.0/5000.0)*sum - 5000

	return poker > 1.03 && poker < 57.4
}
