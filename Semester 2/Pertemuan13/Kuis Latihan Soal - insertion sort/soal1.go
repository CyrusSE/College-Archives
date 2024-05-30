package main

import "fmt"

const NMAX int = 37

type tHimpunan struct {
	anggota [NMAX]int
	panjang int
}

func bacaMasukan(set *tHimpunan) {
	var angka int
	for {
		fmt.Scan(&angka)
		if ada(*set, angka) {
			break
		}
		set.anggota[set.panjang] = angka
		set.panjang++
	}
	urut(set)
}

func ada(set tHimpunan, x int) bool {
	var ada bool
	for i := 0; i < set.panjang; i++ {
		if set.anggota[i] == x {
			ada = true
		}
	}
	return ada
}

func urut(set *tHimpunan) {
	for i := 1; i < set.panjang; i++ {
		save := set.anggota[i]
		j := i
		for j > 0 && save > set.anggota[j-1] {
			set.anggota[j] = set.anggota[j-1]
			j = j - 1
		}
		set.anggota[j] = save
	}
}

func sama(set1, set2 tHimpunan) bool {
	var same bool
	if set1.panjang == set2.panjang {
		if set1.anggota == set2.anggota {
			same = true
		}
	}
	return same
}

func main() {
	var set1, set2 tHimpunan
	fmt.Print("Anggota himpunan 1: ")
	bacaMasukan(&set1)
	fmt.Print("Anggota himpunan 2: ")
	bacaMasukan(&set2)
	fmt.Println("Himpunan 1 = Himpunan 2?", sama(set1, set2))
}
