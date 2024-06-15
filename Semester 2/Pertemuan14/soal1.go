package main

import "fmt"

const NMAX int = 100

type H_A [NMAX]int
type H_B [NMAX]int
type H_C [NMAX]int

func main() {
	var A H_A
	var B H_B
	var C H_C
	var TotalA, TotalB, TotalC int
	A = [100]int{9, 6, 10, 75, 89, 14, 60}
	TotalA = 7
	B = [100]int{5, 14, 35, 10, 75}
	TotalB = 5
	gabungan(A, B, &C, TotalA, TotalB, &TotalC)
	fmt.Print("himpunan C adalah ")
	for i := 0; i < TotalC; i++ {
		fmt.Print(C[i], " ")
	}
	fmt.Println()
	fmt.Println("MAX A :", maxA(A, TotalA))
	fmt.Println("MAX B :", maxB(B, TotalB))
	fmt.Println("MIN C :", minC(C, TotalC))
	findsame(B, A, TotalA, TotalB)
}

func gabungan(A H_A, B H_B, C *H_C, TotalA, TotalB int, TotalC *int) {
	for i := 0; i < TotalA; i++ {
		C[i] = A[i]
	}
	*TotalC = TotalA
	for i := 0; i < TotalB; i++ {
		found := false
		for x := 0; x < *TotalC; x++ {
			if B[i] == C[x] {
				found = true
				break
			}
		}
		if !found {
			C[*TotalC] = B[i]
			*TotalC++
		}
	}
}

func maxA(A H_A, TotalA int) int {
	var MAX int
	for i := 0; i < TotalA; i++ {
		if A[i] > MAX {
			MAX = A[i]
		}
	}
	return MAX
}

func maxB(B H_B, TotalB int) int {
	var MAX int
	for i := 0; i < TotalB; i++ {
		if B[i] > MAX {
			MAX = B[i]
		}
	}
	return MAX
}

func minC(C H_C, TotalC int) int {
	var MIN int
	MIN = C[0]
	for i := 0; i < TotalC; i++ {
		if C[i] < MIN {
			MIN = C[i]
		}
	}
	return MIN
}

func findsame(B H_B, A H_A, TotalA, TotalB int) {
	fmt.Print("Anggota B terdapat pada A : ")
	for i := 0; i < TotalB; i++ {
		for j := 0; j < TotalA; j++ {
			if B[i] == A[j] {
				fmt.Print(B[i], " ")
			}
		}
	}
}
