package main

import "fmt"

func segitiga(a, b, c int) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if segitiga(a, b, c) {
		fmt.Println("segitiga")
	} else {
		fmt.Println("bukan segitiga")
	}
}
