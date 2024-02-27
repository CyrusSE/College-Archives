package main

import "fmt"

func main() {
	var x, n int
	var check bool
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y : (<= n) ")
	fmt.Scan(&n)
	for n > 0 {
		angka := n % 10
		n = n / 10
		if angka == x {
			check = true
		}
	}
	fmt.Print(check)
}
