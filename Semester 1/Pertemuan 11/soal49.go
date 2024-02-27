package main

import "fmt"

func main() {
	var jabatan string
	var masakerja, anak int
	fmt.Print("Masukkan jabatan anda : ")
	fmt.Scan(&jabatan)
	fmt.Print("Masukkan masa kerja : ")
	fmt.Scan(&masakerja)
	fmt.Print("Masukkan jumlah anak : ")
	fmt.Scan(&anak)
	gaji := 0
	tunjangan := 0
	if jabatan == "Staf" {
		if masakerja < 5 {
			gaji = 4000
		} else if masakerja > 10 {
			gaji = 5000
			if anak > 3 {
				tunjangan = 100 * 3
			} else {
				tunjangan = 100 * anak
			}
		} else if masakerja >= 5 && masakerja <= 10 {
			gaji = 4000
			if anak > 3 {
				tunjangan = 100 * 3
			} else {
				tunjangan = 100 * anak
			}
		}
	} else if jabatan == "Manager" || jabatan == "Manajer" {
		if masakerja > 10 {
			gaji = 10000
			if anak > 3 {
				tunjangan = 300 * 3
			} else {
				tunjangan = 300 * anak
			}
		} else if masakerja <= 10 {
			gaji = 8500
			if anak > 3 {
				tunjangan = 300 * 3
			} else {
				tunjangan = 300 * anak
			}
		}
	} else if jabatan == "Direktur" {
		gaji = 20000
		if anak > 3 {
			tunjangan = 500 * 3
		} else {
			tunjangan = 500 * anak
		}
	}
	totalgaji := gaji + tunjangan
	fmt.Println(gaji, "+", tunjangan, "=", totalgaji)
}
