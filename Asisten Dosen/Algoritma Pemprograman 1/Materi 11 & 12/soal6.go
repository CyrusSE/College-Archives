package main

import "fmt"

func main() {
	var t, v, total int
	var check bool
	fmt.Scan(&t)
	for t >= total {
		fmt.Scan(&v)
		total = total + v
		if total >= t {
			check = true
		}
		fmt.Println(check)
	}
}
