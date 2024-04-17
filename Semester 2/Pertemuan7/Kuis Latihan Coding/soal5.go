package main

import "fmt"

const NMax int = 1000

type Baju [NMax]int

func inputB(T, P *Baju, jHari *int) {
	/*I.S.-
	Proses: menerima masukan barang terjual dan pendapatan. Pemasukan berhenti
	apabila barang terjual dan pendapatan bernilai 0. jHari akan terus bertambah
	1 selama barang terjual dan pendapatan tidak bernilai 0.
	F.S. jHari berisi banyaknya hari, Array T (terjual), dan Array P(pendapatan)
	berisi bilangan bulat yang telah diinputkan.*/
	var terjual, pendapatan int
	fmt.Scan(&terjual)
	fmt.Scan(&pendapatan)
	*jHari = 0
	for terjual != 0 && pendapatan != 0 {
		T[*jHari] = terjual
		P[*jHari] = pendapatan
		*jHari = *jHari + 1
		fmt.Scan(&terjual)
		fmt.Scan(&pendapatan)
	}
}

func rata2(T, P Baju, RT, RP *float64, jHari int) {
	/*I.S. Array T berisi barang terjual, Array P berisi pendapatan, dan jHari berisi banyaknya hari.
	Proses: hitung rata-rata barang terjual lalu masukan ke Array RT, dan hitung rata-rata pendapatan
	lalu masukan ke array RP
	F.S. Menyimpan array rata-rata barang terjual (RT) dan pendapatan (RP) selama jhari*/
	var i, totalT, totalP int
	totalT = 0
	totalP = 0
	for i = 0; i < jHari; i++ {
		totalT = totalT + T[i]
		totalP = totalP + P[i]
	}
	*RT = float64(totalT) / float64(jHari)
	*RP = float64(totalP) / float64(jHari)
}

func tampilkan(RT, RP float64, jHari int) {
	/*I.S. Array RT berisi rata-rata barang terjual, array RP berisi rata-rata pendapatan,
	dan jHari berisi banyaknya hari.
	F.S. Menampilkan informasi rata-rata barang terjual dan pendapatan selama beberapa hari.*/
	fmt.Printf("Rata-rata barang yang terjual selama %d hari adalah: %.2f\n", jHari, RT)
	fmt.Printf("Rata-rata pendapatan yang diperoleh selama %d hari adalah: %.2f", jHari, RP)
}

func main() {
	var t, p Baju
	var rata2_T, rata2_P float64
	var jHari int
	inputB(&t, &p, &jHari)
	rata2(t, p, &rata2_T, &rata2_P, jHari)
	tampilkan(rata2_T, rata2_P, jHari)
}
