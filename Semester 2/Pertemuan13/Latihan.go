package main

import "fmt"

const NMAX int = 5

type gacor struct {
	id    int
	nama  string
	nilai int
}

type arr [NMAX]gacor

func main() {
	var ray arr
	ray[0].id = 1
	ray[0].nama = "Clara"
	ray[0].nilai = 70
	ray[1].id = 2
	ray[1].nama = "Mark"
	ray[1].nilai = 90
	ray[2].id = 3
	ray[2].nama = "Ashlyn"
	ray[2].nilai = 100
	ray[3].id = 4
	ray[3].nama = "Jacob"
	ray[3].nilai = 40
	ray[4].id = 5
	ray[4].nama = "Bella"
	ray[4].nilai = 60
	sorting(&ray)
	fmt.Println("Terkecil : ", findmin(ray))
	fmt.Println("Terbesar : ", findmax(ray))
	fmt.Println("Found : ", findbinary(ray, 100), "(Nilai 100)")
}

func sorting(A *arr) {
	for i := 0; i < NMAX; i++ {
		min := i
		for x := i; x < NMAX; x++ {
			if A[min].nilai > A[x].nilai {
				min = x
			}
		}
		A[min], A[i] = A[i], A[min]
	}

}

func findmax(A arr) int {
	save := -1
	for i := 0; i < NMAX; i++ {
		if A[i].nilai > save {
			save = A[i].nilai
		}
	}
	return save
}

func findmin(A arr) int {
	save := A[0].nilai
	for i := 0; i < NMAX; i++ {
		if A[i].nilai < save {
			save = A[i].nilai
		}
	}
	return save
}

func findbinary(A arr, n int) string {
	kiri := 0
	kanan := NMAX - 1
	tengah := (kiri + kanan) / 2
	for kiri < kanan && A[tengah].nilai != n {
		if n < A[tengah].nilai {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
		tengah = (kiri + kanan) / 2
	}
	if A[tengah].nilai != n {
		return "Ganemu"
	} else {
		return A[tengah].nama
	}
}
