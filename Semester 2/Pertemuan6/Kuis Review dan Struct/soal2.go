package main

import "fmt"

func jumlahGanjil(N int, hasil *int) {
	/* I.S. terdefinisi bilangan bulat N
	   F.S. suatu bilangan bulat yang berisi hasil penjumlahan 1 + 3 + 5 + ... + N*/

	for i := 1; i <= N; i += 2 {
		*hasil += i
	}
}

func main() {
	var N, hasil int
	fmt.Scan(&N)
	jumlahGanjil(N, &hasil)
	fmt.Print(hasil)
}
