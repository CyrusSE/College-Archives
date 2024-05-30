package main

import "fmt"

const NMAX int = 1000

type arrayBunga [NMAX]string

func main() {
	var A arrayBunga
	var N int
	fmt.Scan(&N)
	isiArray(&A, N)
	mengurutkan(&A, N)
	tampilArray(A, N)
}

func panjang(bunga string) int {
	return len(bunga)
}

func mengurutkan(tabBunga *arrayBunga, N int) {
	for i := 1; i < N; i++ {
		save := tabBunga[i]
		j := i - 1
		for j >= 0 && panjang(tabBunga[j]) > panjang(save) {
			tabBunga[j+1] = tabBunga[j]
			j = j - 1
		}
		tabBunga[j+1] = save
	}
}

func isiArray(A *arrayBunga, N int) {
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
}

func tampilArray(A arrayBunga, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(A[i])
	}
}
