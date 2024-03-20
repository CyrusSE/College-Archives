package main

import "fmt"

func mengurutkan(a *int, b *int) {
	var temp int
	if *a > *b {
		temp = *a
		*a = *b
		*b = temp
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for a != b {
		mengurutkan(&a, &b)
		fmt.Println(a, b)
		fmt.Scan(&a, &b)
	}
}
