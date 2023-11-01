package fips140

import (
	"testing"
)

func Test_Longruns(t *testing.T) {
	tests := []struct {
		input  []uint64
		streak int
		result bool
	}{
		{generateAll(1), 20000, false},
		{generateAll(0), 20000, false},
	}

	for _, test := range tests {
		set := test.input

		result, streak := LongRun(set)

		if streak != test.streak && result != test.result {
			t.Errorf("Expected:- %d and %t, but got - %d and %t", test.streak, test.result, streak, result)
		}
	}
}

func Test_Sequence(t *testing.T) {
	tests := []struct {
		input  []uint64
		streak int
		result bool
	}{
		{generateAll(1), 20000, false},
		{generateAll(0), 20000, false},
	}

	for _, test := range tests {
		set := test.input

		//set[len(set)-1] = 1

		result, _ := Runs(set)

		if result != test.result {
			t.Errorf("Expected:- %t, but got - %t", true, result)
		}
	}
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
