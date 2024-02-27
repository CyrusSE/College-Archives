package main

import "fmt"

func main() {
	var n int
	var barang1, barang2, barang3, barang4, barang5, check bool
	fmt.Print("Masukkan jumlah anggota tim : ")
	fmt.Scan(&n)
	check = true
	// check2 := true
	for i := 1; i <= n; i++ {
		fmt.Scan(&barang1, &barang2, &barang3, &barang4, &barang5)
		check = check && barang1 && barang2 && barang3 && barang4 && barang5
		// check2 = (!barang1 || !barang2 || !barang3 || !barang4 || !barang5)
		// if !barang1 || !barang2 || !barang3 || !barang4 || !barang5 {
		// 	check = false
		// }
	}
	fmt.Print(check)
}
