package main

import "fmt"

func main() {
	var detik int
	fmt.Print("Masukkan nilai detik : ")
	fmt.Scan(&detik)
	jam := detik / 3600
	detik = detik - (jam * 3600)
	menit := detik / 60
	detik = detik - (menit * 60)
	fmt.Print(jam, " jam ", menit, " menit ", detik, " detik ")
}
