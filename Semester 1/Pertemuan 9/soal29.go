package main

import "fmt"

func main() {
	var k byte
	var check bool
	fmt.Scanf("%c", &k)
	check = (k >= 'A' && k <= 'Z' || k >= 'a' && k <= 'z' || k >= '0' && k <= '9')
	fmt.Println(check)
}
