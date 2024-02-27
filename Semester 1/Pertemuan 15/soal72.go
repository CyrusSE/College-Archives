package main

import "fmt"

func main() {
	var x, angka, angka2 int
	var ascending, descending bool
	fmt.Print("Masukan bilangan bulat positif : ")
	fmt.Scan(&x)
	ascending = true
	descending = true
	for x > 0 {
		angka = x % 10
		x = x / 10
		angka2 = x % 10
		if angka != 0 && angka2 != 0 {
			if angka > angka2 {
				descending = false
			} else if angka < angka2 {
				ascending = false
			}
		}
	}
	if ascending {
		fmt.Println("ascending")
	} else if descending {
		fmt.Println("descending")
	} else if !ascending && !descending {
		fmt.Println("tidak terurut")
	}
}
