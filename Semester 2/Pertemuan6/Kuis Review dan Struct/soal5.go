package main

import "fmt"

type Orang struct {
	nama   string
	gendre string
	alamat string
	agama  string
}

func main() {
	var org Orang

	fmt.Print("Nama : ")
	fmt.Print(&org.nama)

	fmt.Print("Gendre : ")
	fmt.Scan(&org.gendre)

	fmt.Print("Alamat : ")
	fmt.Scan(&org.alamat)

	fmt.Print("Agama : ")
	fmt.Println(&org.agama)

	fmt.Println("Hasil input data : ")
	fmt.Println(org.nama)
	fmt.Println(org.gendre)
	fmt.Println(org.alamat)
	fmt.Println(org.agama)
}
