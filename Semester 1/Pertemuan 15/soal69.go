package main

import "fmt"

func main() {
	var check bool
	var x, total, rata, total2, ratarata, temp float64

	check = true
	temp = x
	ratarata = 0.0
	for check {
		fmt.Scan(&x)
		if x > temp && total != 0 && total != 1 {
			total2 = total2 + 1
		}
		temp = x
		if x < 0 || x > 200 {
			break
		}
		total = total + 1
		rata = rata + x
	}
	if total == 0 && rata == 0 {
		rata = 0
	} else {
		ratarata = rata / total
	}
	fmt.Print(total2, ratarata)
}
