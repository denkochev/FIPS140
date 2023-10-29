package fips140

// currently, test supports only 20 000 bits!
func Monobit(bits_set []uint64) bool {
	var ones int = 0
	var zeros int = 0

	block := 64

	for i := 0; i < len(bits_set); i++ {
		curBlock := bits_set[i]

		if i == len(bits_set)-1 {
			block = 32
		}

		for j := 0; j < block; j++ {
			least_significant_bit := curBlock & 1
			curBlock = curBlock >> 1

			if least_significant_bit == 1 {
				ones += 1
			} else {
				zeros += 1
			}
		}
	}
	// check if set for test equal 20 000 bits
	if ones+zeros == 20000 {
		return 9654 < ones && ones < 10346
	}

	return false
}
