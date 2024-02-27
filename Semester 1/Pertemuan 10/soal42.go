package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Print("Masukkan 4 bilangan jumlah gol : ")
	fmt.Scan(&a, &b, &c, &d)
	terbanyak := 0
	terdikit := 0
	if a >= b && a >= c && a >= d {
		terbanyak = a
	} else if b >= a && b >= c && b >= d {
		terbanyak = b
	} else if c >= a && c >= b && c >= d {
		terbanyak = c
	} else if d >= a && d >= b && d >= c {
		terbanyak = d
	}
	if a <= b && a <= c && a <= d {
		terdikit = a
	} else if b <= a && b <= c && b <= d {
		terdikit = b
	} else if c <= a && c <= b && c <= d {
		terdikit = c
	} else if d <= a && d <= b && d <= c {
		terdikit = d
	}
	fmt.Println(terbanyak, terdikit)
}
