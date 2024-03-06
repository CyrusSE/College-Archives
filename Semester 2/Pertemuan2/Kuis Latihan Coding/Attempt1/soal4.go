package main

import "fmt"

func penjumlahanPecahan(pembilang1, penyebut1, pembilang2, penyebut2 int) int {
	if penyebut1 == penyebut2 {
		return pembilang1 + pembilang2
	}
	return 0
}

func main() {
	var pembilang1, penyebut1, pembilang2, penyebut2 int
	fmt.Scan(&pembilang1, &penyebut1, &pembilang2, &penyebut2)
	fmt.Println(penjumlahanPecahan(pembilang1, penyebut1, pembilang2, penyebut2))
}
