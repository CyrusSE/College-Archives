package main

import "fmt"

func cetakHitungJumlahRata(n, m int, jum *int, rata *float64) {
	/* I.S.: Terdefinisi bilangan bulat n dan m
	   F.S.: jum berisi nilai penjumlahan bilangan bulat dari n hingga m
	         dan rata berisi nilai rata-rata bilangan bulat dari n hingga m */
	var total float64
	for i := n; i <= m; i++ {
		total += 1
		*jum += i
		fmt.Println(i)
	}
	*rata = float64(*jum) / total
}

func main() {
	var n, m, jumlah int
	var rerata float64
	fmt.Scan(&n, &m)
	cetakHitungJumlahRata(n, m, &jumlah, &rerata)
	fmt.Println(jumlah, rerata)
}
