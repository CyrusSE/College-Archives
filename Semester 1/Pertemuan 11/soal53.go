package main

import "fmt"

func main() {
	var jam, menit int
	var jarak float64
	var tarif float64
	fmt.Print("Masukkan waktu pemesanan taksi : (jam & menit) ")
	fmt.Scan(&jam, &menit)
	fmt.Print("Masukkan jarak : ")
	fmt.Scan(&jarak)
	tarif = 0
	if jam >= 5 && jam <= 9 || jam >= 16 && jam <= 19 {
		if jam == 9 || jam == 19 {
			if menit == 00 {
				if jarak > 0 && jarak <= 10 {
					tarif = 5000 * jarak
				} else if jarak > 10 && jarak <= 20 {
					tarif = 4500 * jarak
				}
			}
		} else {
			if jarak > 0 && jarak <= 10 {
				tarif = 5000 * jarak
			} else if jarak > 10 && jarak <= 20 {
				tarif = 4500 * jarak
			}
		}
	} else if jam >= 22 || jam < 5 {
		fmt.Print("Maaf, kami tidak bisa melayani pesanan anda.")
		return
	} else {
		if jarak > 0 && jarak <= 20 {
			tarif = 4000 * jarak
		}
	}
	fmt.Print(tarif)
}
