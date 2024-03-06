package main

import "fmt"

func lowToUpper(kar byte) byte {
	return kar - 32
}

func main() {
	var x byte
	fmt.Scanf("%c", &x)
	this := lowToUpper(x)
	fmt.Printf("%c", this)
}
