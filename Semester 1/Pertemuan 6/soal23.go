package main

import "fmt"

func main() {
	var t, v, total int
	var check bool
	fmt.Print("Masukan nilai T untuk kapasitas tanki : ")
	fmt.Scan(&t)
	for t >= total {
		fmt.Print("Masukkan volume air : ")
		fmt.Scan(&v)
		total = total + v
		if total >= t {
			check = true
		}
		fmt.Println(check)
	}
}
