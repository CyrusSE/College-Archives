package main

import "fmt"

func sortAscending(a, b, c int) {
	/* I.S. terdefinisi berupa 3 bilangan bulat a, b, dan c
	   F.S. a, b, c yang sudah terurut ascending*/
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	fmt.Println(a, b, c)

}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	sortAscending(a, b, c)
}
