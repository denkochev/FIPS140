package main

import (
	"fips140/fips140"
	"fmt"
)

func main() {
	randset := fips140.GetRandomSet(20000)
	monobit := fips140.Monobit(randset)
	fmt.Println(monobit)
}
