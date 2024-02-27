package main

import "fmt"

func main() {
	var karakter int
	var kapital bool
	fmt.Print("Masukkan sebuah karakter: ")
	fmt.Scanf("%c", &karakter)
	kapital = (karakter >= 'A' && karakter <= 'Z')
	fmt.Println(kapital)
}
