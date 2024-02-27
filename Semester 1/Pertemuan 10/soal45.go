package main

import "fmt"

func main() {
	var p int
	var appadfull, appakfull, gow bool
	fmt.Print("Masukkan jumlah orang yang berangkat : ")
	fmt.Scan(&p)
	appad := 3
	naikappad := 0
	appak := 5
	naikappak := 0
	nogo := 0
	i := 0
	i2 := 0
	for p > 0 {
		if i == appad {
			appadfull = true
		}
		if !appadfull {
			i++
			if p >= 5 {
				naikappad += 1
				p -= 5
			} else {
				naikappad += 1
				p -= p
			}
		}
		if i2 == appak {
			appakfull = true
		}
		if appadfull && !appakfull {
			i2++
			if p >= 2 {
				naikappak += 1
				p -= 2
			} else {
				naikappak += 1
				p -= p
			}
		}
		if appadfull && appakfull {
			nogo = p
			p -= p
			gow = true
		}
	}
	if appadfull && appakfull && gow {
		fmt.Print("Dewasa ", naikappad, ", Kecil ", naikappak, ", dan ", nogo, " tak berangkat")
	} else if appadfull && naikappak != 0 {
		fmt.Print("Dewasa ", naikappad, ", Kecil ", naikappak)
	} else {
		fmt.Print("Dewasa ", naikappad)
	}
}
