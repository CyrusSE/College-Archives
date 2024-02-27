package main

import "fmt"

func main() {
	var x, y, z int
	var check bool
	fmt.Print("Masukkan jari jari lingkaran A : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan jari jari lingkaran B : ")
	fmt.Scan(&y)
	fmt.Print("Masukkan jarak titik pusat lingkaran A & B : ")
	fmt.Scan(&z)
	check = (z > x && z > y)
	// if z > x && z > y {
	// 	check = true
	// }
	fmt.Print(check)
}
