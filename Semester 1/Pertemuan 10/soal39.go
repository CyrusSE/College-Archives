package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan 1 bilangan bulat : ")
	fmt.Scan(&x)
	if x < 0 {
		x = -x
	}
	fmt.Println(x)
}
