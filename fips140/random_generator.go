package fips140

import (
	"fmt"
	"math/rand"
)

func GetRandomSet(length int) []uint64 {
	var rand_set []uint64

	if length == 20000 {
		rand_set = make([]uint64, length/64+1)
	} else {
		/*
			future implementation for different sizes
		*/
	}

	block := 64

	for i := 0; i < len(rand_set); i++ {
		var curBlock uint64
		/*
			one block in the set has to be 32 bits instead of 64
			because 313 uint64 == 20 032 bits
			and our tests required 20 000 bits
			we need to get [1]uint32 + [312]uint64
			as a result we get 20 000 random bits,
			for a solution I made last block 32 random bits instead of 64
		*/
		if i == len(rand_set)-1 {
			block = 32
		}

		for j := 1; j <= block; j++ {
			curBlock = curBlock << 1
			curRandBit := uint64(rand.Intn(2))
			curBlock = curBlock | curRandBit
		}
		//fmt.Printf("%064b\n", curBlock)
		rand_set[i] = curBlock
	}

	return rand_set
}

func PrintAllBits(blocks []uint64) string {
	result := ""

	for i := 0; i < len(blocks); i++ {
		if i == len(blocks)-1 {
			result += fmt.Sprintf("%032b", blocks[i])
			//fmt.Printf("%032b\n", blocks[i])
		} else {
			result += fmt.Sprintf("%064b", blocks[i])
			//fmt.Printf("%064b\n", blocks[i])
		}
	}

	return result
}
