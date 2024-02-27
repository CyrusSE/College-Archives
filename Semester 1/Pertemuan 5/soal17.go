package main

import "fmt"

func main() {
	var n int
	var s bool
	fmt.Print("Masukkan nilai n : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			s = true
		} else {
			s = false
		}
		fmt.Println(i, s)
	}
}
