package main

import "fmt"

func main() {
	var arr1 = [8]int{5, 2, 1, 9, 12, 13, 10, 15}
	for i := 1; i < 8; i++ {
		save := arr1[i]
		j := i
		for j > 0 && save > arr1[j-1] {
			arr1[j] = arr1[j-1]
			j = j - 1
		}
		arr1[j] = save
	}
	fmt.Println(arr1)
}
