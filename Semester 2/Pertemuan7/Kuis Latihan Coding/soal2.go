package main

import "fmt"

type pemain [1000]string
type tabInt [1000]int

func main() {
	var n, i int
	var p pemain
	var set1, set2, set3, total tabInt

	//input banyaknya pemain
	fmt.Scan(&n)

	//input nama pemain dan poin yang diperoleh per sesi
	for i = 0; i < n; i++ {
		fmt.Scan(&p[i], &set1[i], &set2[i], &set3[i])
	}

	//menghitung poin dan menampilkan deskripsi. (output kalimat huruf kecil semua)
	for i = 0; i < n; i++ {
		total[i] = set1[i] + set2[i] + set3[i]
		if total[i] < 40 {
			fmt.Println(p[i], "perlu ikut latihan ekstra selama 1 minggu")
		} else if total[i] >= 40 && total[i] < 50 {
			fmt.Println(p[i], "perlu ikut latihan ekstra selama 3 hari")
		} else if total[i] >= 50 {
			fmt.Println(p[i], "tidak perlu mengikuti latihan ekstra")
		}
	}
}
