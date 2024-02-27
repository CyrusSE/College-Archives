package main

import "fmt"

func main() {
	var x, digit1, digit3, digit4, jumlah, tengah int
	var diskon, cashback bool
	fmt.Scan(&x)
	digit1 = (x / 1000)
	digit3 = (x / 10) % 10
	digit4 = x % 10
	jumlah = digit1 + digit3
	tengah = (x / 10) % 100
	diskon = (tengah%2 == 0)
	cashback = (digit4 != 0 && jumlah%digit4 == 0)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
}
