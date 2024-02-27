package main

import "fmt"

func main() {
	var n, m, kpk int
	fmt.Print("Masukkan nilai N dan M : ")
	fmt.Scan(&n, &m)
	for is := 1; kpk == 0; is++ {
		if is%n == 0 && is%m == 0 {
			kpk = is
			fmt.Println(kpk)

		}
	}
}
