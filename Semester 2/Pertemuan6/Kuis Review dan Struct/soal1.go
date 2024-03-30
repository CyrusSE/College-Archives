package main

import "fmt"

func hitungBakteri(x, i, bakteri int) int {
	/*I.S. Terdefinisi nilai bilangan bulat x.
	  /*F.S. fungsi mengembalikan jumlah banyaknya bakteri saat x*/
	for y := i; y <= x; y++ {
		bakteri *= 2
	}
	return bakteri
}

func main() {
	var x, result int
	fmt.Scanln(&x)
	result = hitungBakteri(x, 1, 2)
	fmt.Print(result)
}
