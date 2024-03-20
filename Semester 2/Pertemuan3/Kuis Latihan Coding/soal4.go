package main

import "fmt"

func totalPelajar(a string, b int, totSD, totSMP, totSMA *int) {
	/* I.S. terdefisi pelajar dan jumlah
	   F.S. totSD adalah jumlah pelajar SD, totSMP adalah jumlah pelajar SMP, totSMA adalah jumlah pelajar SMA*/
	   
		...
}

func main() {
	var jum, totSD, totSMP, totSMA int
	var jenis string
	fmt.Scan(&jenis, &jum)
	for jenis != "x" || jum != 0 {
		totalPelajar(jenis, jum, &totSD, &totSMP, &totSMA)
		fmt.Scan(&jenis, &jum)
	}
	fmt.Println(totSD, totSMP, totSMA)
}
