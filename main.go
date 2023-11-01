package main

import (
	"fips140/fips140"
	"fmt"
)

func main() {
	randset := fips140.GetRandomSet(20000)
	stringBits := fips140.PrintAllBits(randset)
	fmt.Println("rangom bits -> ", len(stringBits))

	monobit := fips140.Monobit(randset)
	fmt.Println("The Monobit Test -> ", monobit)

	longrun, streak := fips140.LongRun(randset)
	fmt.Println("The Long Run Test -> ", longrun, "| streak for this set = ", streak)

	pokerTest := fips140.Poker(randset)
	fmt.Println("The Poker Test -> ", pokerTest)

	longsequence, curSequenceTable := fips140.Runs(randset)
	fmt.Println("The Runs Test -> ", longsequence, "| table for this set = ", curSequenceTable)

	fmt.Println("---------------------------------------------------------------------------------------------------")

}
