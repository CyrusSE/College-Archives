package main

import "fmt"

func main() {
	var member string
	var a, b, c, d, e int
	var cashback, diskon, tambahan float64
	fmt.Scan(&member, &a, &b, &c, &d, &e)
	tambahan = 0.0
	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = 3.1 * (float64(a) + float64(b) + float64(c))
		if member == "Yes" {
			tambahan = cashback * (15.0 / 100)
			cashback = cashback + tambahan
		}
		if cashback > 35 {
			cashback = 35
		}
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
		diskon = 1.7 * (float64(c) + float64(d) + float64(e))
		if member == "Yes" {
			tambahan = diskon * (15.0 / 100)
			diskon = diskon + tambahan
		}
		if diskon > 50 {
			diskon = 50
		}
	} else {
		diskon = 1.7 * (float64(c) + float64(d) + float64(e))
		cashback = 3.1 * (float64(a) + float64(b) + float64(c))
		if member == "Yes" {
			tambahan = cashback * (15.0 / 100)
			cashback = cashback + tambahan
			tambahan = diskon * (15.0 / 100)
			diskon = diskon + tambahan
		}
		if cashback > 35 {
			cashback = 35
		}
		if diskon > 50 {
			diskon = 50
		}
	}
	fmt.Print("Cashback: ", cashback, "% Diskon: ", float32(diskon), "%")
}
