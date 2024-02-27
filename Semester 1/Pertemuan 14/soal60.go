package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tipea, tipeb, tipec int
	fmt.Print("Masukkan beberapa karakter (A/B/C) : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	for _, gg := range input {
		if gg == 65 {
			tipea += 1
		} else if gg == 66 {
			tipeb += 1
		} else if gg == 67 {
			tipec += 1
		}
	}
	fmt.Println("Tipe A = ", tipea)
	fmt.Println("Tipe B = ", tipeb)
	fmt.Println("Tipe C = ", tipec)
}
