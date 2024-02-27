package main

import "fmt"

func main() {
	var nim, x int
	var check bool
	fmt.Scan(&x)
	nim = 13065841
	for nim > 10 {
		nim = nim - 10
	}
	nim = nim + 125
	for nim > 0 {
		nim = nim - x
	}
	check = (nim == 0)
	fmt.Print(check)
}
