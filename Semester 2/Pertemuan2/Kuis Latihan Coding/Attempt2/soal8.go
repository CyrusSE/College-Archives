package main

import "fmt"

func bukanNol(x int) bool {
	return x != 0
}

func main() {
	var bilangan int
	var jenis bool
	fmt.Scan(&bilangan)
	jenis = bukanNol(bilangan)
	fmt.Println(jenis)
}
