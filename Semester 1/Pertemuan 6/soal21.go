package main

import "fmt"

func main() {
	var gula, kopi, gula2, kopi2 int
	var bisagula, bisakopi int
	fmt.Print("Masukkan banyak gula tersedia : ")
	fmt.Scan(&gula)
	fmt.Print("Masukkan banyak kopi tersedia : ")
	fmt.Scan(&kopi)
	fmt.Print("Masukkan perlu banyak gula : ")
	fmt.Scan(&gula2)
	for gula2 > gula {
		fmt.Println("Banyak gula harus lebih besar daripada keperluan gula")
		fmt.Print("Masukkan perlu banyak gula : ")
		fmt.Scan(&gula2)
	}
	fmt.Print("Masukkan perlu banyak kopi : ")
	fmt.Scan(&kopi2)
	for kopi2 > kopi {
		fmt.Println("Banyak kopi harus lebih besar daripada keperluan kopi")
		fmt.Print("Masukkan perlu banyak kopi : ")
		fmt.Scan(&kopi2)
	}
	bisagula = gula / gula2
	bisakopi = kopi / kopi2
	if bisagula >= bisakopi {
		fmt.Print(bisakopi)
	} else if bisagula <= bisakopi {
		fmt.Print(bisagula)
	}
}
