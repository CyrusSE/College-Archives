package main

import "fmt"

func main() {
	var huruf byte

	fmt.Scanf("%c", &huruf)

	if huruf == 'A' || huruf == 'I' || huruf == 'U' || huruf == 'E' || huruf == 'O' || huruf == 'a' || huruf == 'i' || huruf == 'u' || huruf == 'e' || huruf == 'o' {
		fmt.Print("Bukan konsonan")
	} else {
		fmt.Print("Konsonan")
	}
}
