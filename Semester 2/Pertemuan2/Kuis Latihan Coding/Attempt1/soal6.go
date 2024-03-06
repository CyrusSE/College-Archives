package main

import "fmt"

func posisiBenda(x float64) float64 {
	return (x * x * x) - 12*(x*x) + 36*x - 30
}

func main() {
	var waktu, posisi float64
	fmt.Scan(&waktu)
	posisi = posisiBenda(waktu)
	fmt.Println(posisi)
}
