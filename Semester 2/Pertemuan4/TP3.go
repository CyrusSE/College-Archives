package main

import "fmt"

func cetak_rerata(END, start, jumlah int, rata float64) {
	if start > END {
		return
	} else {
		jumlah += start
		rata = float64(jumlah) / float64(start)
		fmt.Println(jumlah, rata)
		cetak_rerata(END, start+1, jumlah, rata)
	}
}

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata(N, 1, 0, 0)
}
