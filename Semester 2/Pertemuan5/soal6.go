package main

import "fmt"

func genap(x, start int) {
	if start > x {
		fmt.Print()
	} else {
		if start%2 == 0 {
			fmt.Println(start)
		}
		genap(x, start+1)
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	genap(x, 1)
}
