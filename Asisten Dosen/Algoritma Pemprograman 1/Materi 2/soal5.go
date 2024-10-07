package main

import "fmt"

func main() {
	var a, b, c int
	var u1, u2, u3 float64
	fmt.Scan(&a, &b, &c)
	u1 = float64(a) + float64(a)*0.05
	u2 = float64(b) + float64(b)*0.05
	u3 = float64(c) + float64(c)*0.05
	fmt.Println(u1, u2, u3)
}
