package main

import "fmt"

func faktor(n, start int) {
	if start == n {
		fmt.Println(start)
	} else {
		if n%start == 0 {
			fmt.Println(start)
		}
		faktor(n, start+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	faktor(n, 1)
}
