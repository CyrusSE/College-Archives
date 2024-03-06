package main

import "fmt"

func hurufVokal(x byte) bool {
	return x == 'a' || x == 'A' || x == 'i' || x == 'I' || x == 'u' || x == 'U' || x == 'e' || x == 'E' || x == 'o' || x == 'O'
}

func main() {
	var huruf byte
	var vokal bool
	fmt.Scanf("%c", &huruf)
	vokal = hurufVokal(huruf)
	fmt.Println(vokal)
}
