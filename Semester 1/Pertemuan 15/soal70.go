package main

import "fmt"

func main() {
	var x, total int
	for i := 1; i <= 5; i++ {
		fmt.Print("Hari ", i, " : ")
		fmt.Scan(&x)
		for x < 0 || x > 200 {
			fmt.Print("Hari ", i, " : ")
			fmt.Scan(&x)
		}
		total = total + x
	}
	fmt.Print("Jumlah pengunjung : ", total, " orang")
}
