package main

import "fmt"

func main() {
	var gula, kopi, perlugula, perlukopi int
	var bisagula, bisakopi int
	fmt.Scan(&gula)
	fmt.Scan(&kopi)
	fmt.Scan(&perlugula)
	for perlugula > gula {
		fmt.Scan(&perlugula)
	}
	fmt.Scan(&perlukopi)
	for perlukopi > kopi {
		fmt.Scan(&perlukopi)
	}
	bisagula = gula / perlugula
	bisakopi = kopi / perlukopi
	if bisagula >= bisakopi {
		fmt.Print(bisakopi)
	} else if bisagula <= bisakopi {
		fmt.Print(bisagula)
	}
}
