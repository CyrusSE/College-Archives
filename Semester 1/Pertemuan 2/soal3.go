package main

import "fmt"

func main() {
	var x, d1, d2, d3 int
	fmt.Print("Masukkan bilangan bulat positif x (<= 999): ")
	fmt.Scan(&x)
	if x < 0 || x > 999 {
		fmt.Println("Bilangan harus positif dan <= 999")
		return
	}

	d1 = x / 100
	d2 = (x / 10) % 10
	d3 = x % 10

	fmt.Println("Keluaran :", d1, d2, d3)
	fmt.Println("x = ", x, ", maka d1 = ", d1, ", d2 = ", d2, ", dan d3 = ", d3)
}
