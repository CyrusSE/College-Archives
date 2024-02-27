package main

import "fmt"

func main() {
	var x float32
	fmt.Print("Masukan masa tubuh manusia : ")
	fmt.Scan(&x)
	merkurius := 3.72 * x
	venus := 8.91 * x
	bumi := 9.8 * x
	mars := 3.72 * x
	yupiter := 23.22 * x
	saturnus := 9.01 * x
	uranus := 8.72 * x
	neptunus := 11.07 * x
	fmt.Println(merkurius, venus, bumi, mars, yupiter, saturnus, uranus, neptunus)
}
