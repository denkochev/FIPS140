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

}
