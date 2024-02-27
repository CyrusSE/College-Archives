package main

import "fmt"

func main() {
	var jarijari, LuasP float64
	fmt.Print("Masukan angka untuk jari jari : ")
	fmt.Scan(&jarijari)
	LuasP = 4 * (22.0 / 7) * (jarijari * jarijari)
	fmt.Println(LuasP)
}
