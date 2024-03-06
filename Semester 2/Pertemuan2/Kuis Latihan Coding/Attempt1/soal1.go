package main

import "fmt"

func kelilingPersegi(x int) int {
	return x * 4
}

func main() {
	var sisi int
	fmt.Scan(&sisi)
	fmt.Println(kelilingPersegi(sisi))
}
