package main

import "fmt"

func main() {
	var arr1 = [8]int{5, 2, 1, 9, 12, 13, 10, 15}
	for i := 0; i < 8; i++ {
		min := i
		for x := i + 1; x < 8; x++ {
			if arr1[min] > arr1[x] {
				min = x
			}
		}
		arr1[i], arr1[min] = arr1[min], arr1[i]
	}
	fmt.Println(arr1)
}

//Selection sort
