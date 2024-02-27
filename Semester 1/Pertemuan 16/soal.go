package main

import "fmt"

func main() {
	var n, r, hasiln, hasilr, hasilp, total int
	fmt.Scan(&n, &r)
	hasiln = n
	hasilr = r
	hasilp = n - r
	for i := n - 1; i >= 1; i-- {
		hasiln = hasiln * i
	}
	for i := r - 1; i >= 1; i-- {
		hasilr = hasilr * i
	}
	for i := hasilp - 1; i >= 1; i-- {
		hasilp = hasilp * i
	}
	total = hasiln / (hasilr * hasilp)
	fmt.Println(total)
}
