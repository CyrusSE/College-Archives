package main

import "fmt"

func main() {
	var x, total int
	var check bool
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			total = total + 1
		}
	}
	if total == 2 {
		check = true
	}
	fmt.Print(check)
}
