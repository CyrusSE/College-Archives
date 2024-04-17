package main

import "fmt"

const NMax int = 1000

type tabInt [NMax]int
type tabStr [NMax]string

func main() {
	var nama tabStr
	var jam, menit, detik, total tabInt
	var n, i, min, idx int

	//input banyaknya pelari
	fmt.Scan(&n)

	//input nama, jam, menit, detik
	for i = 0; i < n; i++ {
		fmt.Scan(&nama[i], &jam[i], &menit[i], &detik[i])
	}

	//konversi ke detik
	for i = 0; i < n; i++ {
		total[i] = jam[i]*3600 + menit[i]*60 + detik[i]
	}

	//mencari peserta yang mencapai garis finish tercepat (mencari nilai minimum)
	min = total[0]
	idx = 0
	for i = 1; i < n; i++ {
		if total[i] < min {
			min = total[i]
			idx = i
		}
	}

	//menampilkan informasi peserta yang mencapai garis finish tercepat
	fmt.Println("Peserta tercepat adalah budi dengan waktu", jam[idx], "jam", menit[idx], "menit", detik[idx], "detik")
}
