package main

import "fmt"

const NMax = 1000

var bil []int

func jumlahBilangan(bilangan []int) int {
	var jumlah int
	for i := 0; i < len(bilangan); i++ {
		jumlah += bilangan[i]
	}
	return jumlah
}

func main() {
	var angka, total int

	for {
		fmt.Scan(&angka)
		if angka < 0 {
			break
		}
		bil = append(bil, angka)
	}
	total = jumlahBilangan(bil)
	fmt.Println(total)
}
