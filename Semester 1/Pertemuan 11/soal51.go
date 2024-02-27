package main

import "fmt"

func main() {
	var n int
	var uang int
	fmt.Print("Masukkan berapa banyak minggu : ")
	fmt.Scan(&n)
	totaluang := 0
	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan banyak uang minggu ke ", i, " : ")
		fmt.Scan(&uang)
		totaluang = uang + totaluang
	}
	fmt.Println(totaluang)
}
