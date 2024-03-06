package main

import "fmt"

func persamaan(x float64) float64 {
	return (x*x + 1/x) * (x*x + 1/x)
}

func main() {
	var bilangan float64
	fmt.Scan(&bilangan)
	fmt.Println(persamaan(bilangan))

}
