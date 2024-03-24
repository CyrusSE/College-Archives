package main

import "fmt"

func jumlahHari(bulan int, jmlHari *int) {
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		*jmlHari = 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		*jmlHari = 30
	} else if bulan == 2 {
		*jmlHari = 28
	} else {
		*jmlHari = 0
	}
}

func main() {
	var x, bulan, hari int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		fmt.Scan(&bulan)
		jumlahHari(bulan, &hari)
		fmt.Println(hari)
	}
}
