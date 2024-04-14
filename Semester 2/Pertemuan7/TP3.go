package main

import "fmt"

const NMax int = 20

type tabChar [NMax]rune

func main() {
	var kata tabChar
	var nChar int
	fmt.Scan(&nChar)
	baca(&kata, &nChar)
	cetak(kata, nChar)
}

func baca(k *tabChar, n *int) {
	/*
		IS: Array karakter (k) dan banyak elemen (n) terdefinisi sembarang
		FS: Array karakter (k) dan banyak elemen (n) terdefinisi
	*/
	if *n > NMax {
		*n = NMax
	}
	for i := 0; i < *n; i++ {
		fmt.Scanf("%c", &k[i])
	}
}

func cetak(k tabChar, n int) {
	/*
		IS: Array k dan n terdefinisi
		FS: Array k tercetak di layar dengan urutan terbalik
	*/
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%c", k[i])
	}
}
