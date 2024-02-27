package main

import "fmt"

func main() {
	var posisi int
	var ditarik, status_gas bool
	// fmt.Print("Masukkan posisi gigi : ")
	fmt.Scan(&posisi, &ditarik, &status_gas)
	// fmt.Print("Masukkan status kopling : ")
	// fmt.Scan(&ditarik)
	// fmt.Print("Masukkan status gas : ")
	// fmt.Scan(&status_gas)
	status := ""
	if posisi >= 0 {
		if posisi == 0 && ditarik || !ditarik && status_gas {
			status = "Mesin menyala dan motor tidak berjalan"
		} else if posisi == 0 && ditarik && !status_gas {
			status = "Mesin mati"
		} else if !ditarik && status_gas {
			status = "Motor Berjalan"
		} else if ditarik && !status_gas || ditarik && status_gas {
			status = "Mesin menyala dan motor tidak berjalan"
		} else if !ditarik && !status_gas {
			status = "Mesin mati"
		}

	}
	fmt.Print(status)
}
