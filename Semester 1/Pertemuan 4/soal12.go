package main

import "fmt"

func main() {
	var suku1, suku2 int
	fmt.Print("Masukan angka untuk suku1 : ")
	fmt.Scan(&suku1)
	fmt.Print("Masukan angka untuk suku2 : ")
	fmt.Scan(&suku2)
	suku3 := suku1 + suku2
	suku4 := suku3 + suku2
	suku5 := suku3 + suku4
	fmt.Println(suku3, suku4, suku5)
}
