package main

import "fmt"

func main() {
	var quantum, galactum, nebula, stellaris, lumin int
	fmt.Scanln(&lumin)
	for lumin > 1 {
		lumin -= 20
		stellaris += 1
		if stellaris >= 2 && (quantum == 0 || galactum == 0 || nebula == 0) {
			nebula += 1
			stellaris -= 2
		}
		if nebula >= 3 {
			galactum += 1
			nebula -= 3
		}
		if galactum >= 10 && quantum == 0 {
			quantum += 1
			galactum -= 10
		}
	}
	fmt.Println(quantum, galactum, nebula, stellaris, lumin)
}
