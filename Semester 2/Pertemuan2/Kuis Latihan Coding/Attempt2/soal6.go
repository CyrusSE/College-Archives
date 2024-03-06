package main

import "fmt"

func numToTeks(x int) string {
	var teks string
	if x > 0 {
		teks = "positif"
	} else if x < 0 {
		teks = "negatif"
	} else {
		teks = "netral"
	}
	return teks
}

func main() {
	var bilangan int
	var hasil string
	fmt.Scan(&bilangan)
	hasil = numToTeks(bilangan)
	fmt.Println(hasil)

}
