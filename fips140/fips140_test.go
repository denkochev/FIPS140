package fips140

import (
	"math/rand"
	"testing"
)

// passed
func Test_Monobit(t *testing.T) {
	tests := []struct {
		input  []uint64
		result bool
	}{
		{generateAll(1), false},
		{generateAll(0), false},
		{generateOnes(9654), false},
		{generateOnes(9653), false},
		{generateOnes(9655), true},
		{generateOnes(10346), false},
		{generateOnes(10345), true},
	}

	for idx, test := range tests {
		set := test.input

		result := Monobit(set)

		if result != test.result {
			t.Errorf("For test %d Expected:- %t, but got - %t", idx, test.result, result)
		}
	}
}

// passed
func Test_Longrun(t *testing.T) {
	tests := []struct {
		input  []uint64
		streak int
		result bool
	}{
		{generateAll(1), 20000, false},
		{generateAll(0), 20000, false},
		{generateStreak(20), 20, true},
		{generateStreak(30), 30, true},
		{generateStreak(34), 34, false},
		{generateStreak(33), 33, true},
		{generateStreak(35), 33, false},
	}

	for _, test := range tests {
		set := test.input

		result, streak := LongRun(set)

		if streak != test.streak && result != test.result {
			t.Errorf("Expected:- %d and %t, but got - %d and %t", test.streak, test.result, streak, result)
		}
	}
}

// passed
func Test_Poker(t *testing.T) {
	tests := []struct {
		input  []uint64
		result bool
	}{
		{generateAll(1), false},
		{generateAll(0), false},
		{generateOnes(9654), false},
		{generateOnes(9653), false},
		{generateOnes(9655), false},
		{generateOnes(10346), false},
		{generateOnes(10345), false},
		{GetRandomSet(20000), true},
	}

	for idx, test := range tests {
		set := test.input

		result := Poker(set)

		if result != test.result {
			t.Errorf("For test %d Expected:- %t, but got - %t", idx, test.result, result)
		}
	}
}

// passed
func Test_Runs(t *testing.T) {
	tests := []struct {
		input  []uint64
		result bool
	}{
		{generateAll(1), false},
		{generateAll(0), false},
		{generateOnes(9654), false},
		{generateOnes(9653), false},
		{generateOnes(9655), false},
		{generateOnes(10346), false},
		{generateOnes(10345), false},
		{GetRandomSet(20000), true},
	}

	for _, test := range tests {
		set := test.input

		//set[len(set)-1] = 1

		result, _ := Runs(set)

		if result != test.result {
			t.Errorf("Expected:- %t, but got - %t", test.result, result)
		}
	}
}

// generator for fixed amount of 1 bits
func generateOnes(amount int) []uint64 {
	rand_set := make([]uint64, 313)

	block := 64

	for i := 0; i < len(rand_set); i++ {

		var curBlock uint64

		if i == len(rand_set)-1 {
			block = 32
		}

		for j := 1; j <= block; j++ {
			var bit uint64 = 1
			if amount <= 0 {
				bit = 0
			}
			amount -= 1

			curBlock = curBlock << 1
			curBlock = curBlock | bit
		}
		//fmt.Printf("%064b\n", curBlock)
		rand_set[i] = curBlock
	}

	return rand_set
}

// generator for same bit 20 000 times
func generateAll(bit uint64) []uint64 {
	rand_set := make([]uint64, 313)

	block := 64

	for i := 0; i < len(rand_set); i++ {

		var curBlock uint64

		if i == len(rand_set)-1 {
			block = 32
		}

		for j := 1; j <= block; j++ {
			curBlock = curBlock << 1
			curBlock = curBlock | bit
		}
		//fmt.Printf("%064b\n", curBlock)
		rand_set[i] = curBlock
	}

	return rand_set
}

// MAX 64 SAME BIT STREAK
func generateStreak(streak int) []uint64 {
	set := make([]uint64, 313)

	block := 64
	// генеруємо однакову послідовність
	var firstBlock uint64 = 0
	for j := 1; j <= block; j++ {
		var bit uint64 = 1
		// only streak amount of bits should be same
		if j > streak+1 {
			bit = uint64(rand.Intn(2))
		} else if j > streak {
			bit = 0
		}

		firstBlock = firstBlock << 1
		firstBlock = firstBlock | bit
	}
	set[0] = firstBlock

	for i := 1; i < len(set); i++ {

		var curBlock uint64

		if i == len(set)-1 {
			block = 32
		}

		for j := 1; j <= block; j++ {
			curBlock = curBlock << 1
			curBlock = curBlock | uint64(rand.Intn(2))
		}
		set[i] = curBlock
	}
	// make sure that our streak is not bigger because of random
	// Встановлення MSB у 0
	set[1] = set[1] & ^(uint64(1) << 63)

	return set
}
