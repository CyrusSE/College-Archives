package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	*l = 3.14 * r * r
	*k = 2 * 3.14 * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, toLuas, totKel *float64) {
	*toLuas = lL + lP
	*totKel = kL + kP
}

func main() {
	var R, S, LL, LP, KL, KP, TL, TP float64
	fmt.Scan(&R, &S)
	if R != 0 && S != 0 {
		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
	}
	for R != 0 && S != 0 {
		hitungLuasKelilingLingkaran(R, &LL, &KL)
		hitungLuasKelilingPersegi(S, &LP, &KP)
		hitungTotal(LL, LP, KL, KP, &TL, &TP)
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", R, S, LL, LP, KL, KP, TL, TP)
		fmt.Scan(&R, &S)
	}
}
