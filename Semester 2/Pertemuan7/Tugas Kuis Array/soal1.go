package main

import "fmt"

const NMax int = 256

type arr [NMax]int

func isiArray(array *arr, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&array[i])
	}
}

func reverseArray(array *arr, n int) {
	for i := 0; i < n/2; i++ {
		array[i], array[n-i-1] = array[n-i-1], array[i]
	}
}

func palindrom(array arr, n int) bool {
	var check bool = true
	for i := 0; i < n/2; i++ {
		if array[i] != array[n-i-1] {
			check = false
		}
	}
	return check
}

func main() {
	var array arr
	var n int
	isiArray(&array, &n)
	reverseArray(&array, n)
	if palindrom(array, n) == false {
		fmt.Println("Array tidak memiliki pola palindrom")
	} else {
		fmt.Println("Array memiliki pola palindrom")
	}
}
