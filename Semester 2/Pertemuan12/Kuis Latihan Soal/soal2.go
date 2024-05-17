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
	for i := 0; i < N; i++ {
		min := i
		for x := i; x < N; x++ {
			if panjang(tabBunga[min]) > panjang(tabBunga[x]) {
				min = x
			}
		}
		tabBunga[i], tabBunga[min] = tabBunga[min], tabBunga[i]
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
