package main

import "fmt"

func main() {
	var bilangan int
	var angka1, angka2, hasil int
	fmt.Print("Masukan bilangan : ")
	fmt.Scan(&bilangan)
	angka1 = bilangan / 10
	angka2 = bilangan % 10
	hasil = angka1*1000 + angka1*100 + angka2*10 + angka2
	fmt.Println(hasil)
}
