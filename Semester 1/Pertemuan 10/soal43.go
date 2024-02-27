package main

import "fmt"

func main() {
	var h1, m1, h2, m2 int
	fmt.Print("Masukkan 4 nilai : ")
	fmt.Scan(&h1, &m1, &h2, &m2)
	jam := h2 - h1
	if h1 > h2 {
		jam = (h2 + 12) - h1
	}
	menit := m2 - m1
	if m1 > m2 {
		jam = jam - 1
		menit = 60 - (m1 - m2)
	}
	if h1 == h2 && m1 > m2 {
		jam = jam + 12
	}
	fmt.Println(jam, menit)

}
