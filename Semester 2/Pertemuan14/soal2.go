package main

import "fmt"

const NMAX int = 100

type data struct {
	nim           int
	nama          string
	nilai_uts     int
	nilai_uas     int
	nilai_lainnya int
	nilai_akhir   float64
}

type mahasiswa [NMAX]data

func main() {
	var mhs mahasiswa
	var nData int
	inputArray(&mhs, &nData)
	selection_search(&mhs, nData)
	fmt.Println()
	fmt.Println("selection Search")
	printArray(mhs, nData)
	fmt.Println()
	insertion_search(&mhs, nData)
	fmt.Println("insertion Search")
	printArray(mhs, nData)

}

func printArray(A mahasiswa, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(A[i].nim, A[i].nama, A[i].nilai_uts, A[i].nilai_uas, A[i].nilai_lainnya, A[i].nilai_akhir)
	}
}

func inputArray(A *mahasiswa, N *int) {
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Scan(&A[i].nim, &A[i].nama, &A[i].nilai_uts, &A[i].nilai_uas, &A[i].nilai_lainnya)
		A[i].nilai_akhir = ((float64(A[i].nilai_uts) * 0.3) + (float64(A[i].nilai_uas) * 0.45) + (float64(A[i].nilai_lainnya) * 0.25))
	}
}

func selection_search(A *mahasiswa, N int) {
	for i := 0; i < N; i++ {
		min := i
		for x := i; x < N; x++ {
			if A[min].nilai_akhir > A[x].nilai_akhir {
				min = x
			}
		}
		A[min], A[i] = A[i], A[min]
	}
}

func insertion_search(A *mahasiswa, N int) {
	for i := 1; i < N; i++ {
		save := A[i]
		j := i
		for j > 0 && save.nilai_akhir > A[j-1].nilai_akhir {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = save
	}
}
