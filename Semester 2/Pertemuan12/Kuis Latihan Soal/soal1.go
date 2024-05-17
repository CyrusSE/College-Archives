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
	SelectionSort(&arr, x)
	fmt.Println(median(arr, x))
}

func SelectionSort(A *tabInt, T int) {
	for i := 0; i < T; i++ {
		minIdx := i
		for j := i; j < T; j++ {
			if A[j] > A[minIdx] {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}
}

func median(A tabInt, T int) float64 {
	if T%2 == 0 {
		return float64(A[T/2-1]+A[T/2]) / 2.0
	}
	return float64(A[T/2])
}
