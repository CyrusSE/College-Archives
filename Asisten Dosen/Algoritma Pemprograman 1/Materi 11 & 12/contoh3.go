package main

import "fmt"

func main() {
	var i int
	i = 1
	for {
		fmt.Println("Perulangan ke", i)
		i++
		if i > 5 {
			break
		}
	}
}

program perulangan
kamus
	i : integer
algoritma
	i <- 1
	repeat
		output("Perulangan ke", i)
		i <- i + 1
	until i > 5
endprogram
