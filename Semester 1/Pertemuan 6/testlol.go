package main

import "fmt"

func main() {
	var x int
	var d0, d1 int
	var check, check2 bool
	fmt.Scan(&x)
	check = true
	for check {
		d0 = x % 10
		d1 = (x / 10) % 10
		x = x / 10
		check2 = d0-d1 == 1 || d1-d0 == 1
		check = check2 && x/10 > 0
	}
	fmt.Println(check2)
}
