package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var x int
	var arr tabInt
	fmt.Scan(&x)
	for i := 0; i < x; i++ {
		fmt.Scan(&arr[i])
	}
	InsertionSort(&arr, x)
	fmt.Println(median(arr, x))
}

func InsertionSort(A *tabInt, T int) {
	for i := 1; i < T; i++ {
		save := A[i]
		j := i - 1
		for j >= 0 && A[j] > save {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = save
	}
}

func median(A tabInt, T int) float64 {
	if T%2 == 0 {
		return float64(A[T/2-1]+A[T/2]) / 2.0
	}
	return float64(A[T/2])
}
