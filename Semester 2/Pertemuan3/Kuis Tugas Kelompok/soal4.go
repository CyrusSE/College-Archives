package main

import "fmt"

var a int

func proB(b, c *int) {
	*b = *b + *c
	*c = a + *b + *c
}

func main() {
	a = 10
	proB(&a, &a)
	fmt.Println(a)
}

//A = 10
//B = 40
//C = 20
//D = 60
