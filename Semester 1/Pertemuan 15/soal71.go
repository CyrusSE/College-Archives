package main

import "fmt"

func main() {
	var a, b int
	var input, pemenang string
	var check bool
	for i := 1; i <= 10; i++ {
		fmt.Scan(&input)
		if input == "A" {
			a = a + 1
		} else if input == "B" {
			b = b + 1
		}
		if !check {
			if a == 4 {
				pemenang = "A"
				check = true
			} else if b == 4 {
				pemenang = "B"
				check = true
			}
		}
	}
	if a > 5 || b > 5 {
		fmt.Print("tidak valid")
	} else {
		fmt.Print("Pemenang : ", pemenang)
	}
}
