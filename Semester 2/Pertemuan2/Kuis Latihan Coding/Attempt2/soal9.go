package main

import "fmt"

func lowToUpper(x byte) byte {
	return x - 32
}

func main() {
	var alfaLow, alfaUp byte
	fmt.Scanf("%c", &alfaLow)
	alfaUp = lowToUpper(alfaLow)
	fmt.Printf("%c", alfaUp)
}
