package main

import "fmt"

func main() {
	var x byte
	var check bool
	fmt.Scanf("%c", &x)
	check = x >= 'A' && x <= 'Z'
	fmt.Println(check)
}
