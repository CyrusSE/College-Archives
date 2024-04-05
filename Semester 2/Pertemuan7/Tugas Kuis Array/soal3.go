package main

import "fmt"

type array [100]int

func isiArray(n *int, arr *array) {
	/*I.S. Terdefinisi pointer n bilangan bulat dan pointer array untuk menampung bilangan bulat
	  F.S. n terisi jumlah data yang diinput dan pointer array terisi bilangan bulat sebanyak n*/
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(&arr[i])
	}
}

func median(n int, arr array) float64 {
	/* mengembalikan median atau nilai tengah dari array */
	return (float64(arr[(n-1)/2]) + float64(arr[(n/2)])) / 2
}

func main() {
	var arr array
	var n int
	isiArray(&n, &arr)
	fmt.Print("Median : ", median(n, arr))
}
