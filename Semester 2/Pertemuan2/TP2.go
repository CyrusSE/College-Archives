package main

import "fmt"

func reamur(C float64) float64 {
	return (4.0 / 5.0) * C
}

func fahrenheit(C float64) float64 {
	return ((9.0 / 5.0) * C) + 32
}

func kelvin(C float64) float64 {
	return 273.0 + C
}

func main() {
	var awal, akhir, step float64
	var C, R, F, K float64
	fmt.Scan(&awal, &akhir, &step)

	fmt.Printf("%10s %10s %10s %10s\n", "Celsius", "Reamur", "Fahrenheit", "Kelvin")

	for C = awal; C <= akhir; C += step {
		R = reamur(C)
		F = fahrenheit(C)
		K = kelvin(C)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", C, R, F, K)
	}

}
