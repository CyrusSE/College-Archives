package main

import "fmt"

func main() {
	var uang int
	var dompet int
	fmt.Print("Masukkan uang : ")
	fmt.Scan(&uang)
	for uang != 0 {
		if uang >= 0 {
			dompet = dompet + uang
		} else {
			dompet = dompet - (-uang)
		}
		fmt.Print("Masukkan uang : ")
		fmt.Scan(&uang)
	}
	fmt.Print("Total uang di dompet : ", dompet)
}
