package main

import "fmt"

func main() {
	var x, y float64
	fmt.Print("Masukkan 2 keuntungan bulan ini dan bulan sebelumnya : ")
	fmt.Scan(&x, &y)
	if x < y {
		untung := y - x
		fmt.Print("Peningkatan sebesar ", untung)
	} else if x > y {
		rugi := x - y
		fmt.Print("Penurunan sebesar ", rugi)
	} else if x == y {
		fmt.Print("Tetap")
	}
}
