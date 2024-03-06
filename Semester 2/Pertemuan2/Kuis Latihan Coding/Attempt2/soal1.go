package main

import "fmt"

func konversiMeterMil(x float64) float64 {
	return x * 0.00062137119
}

func main() {
	var meter, mil float64
	fmt.Scan(&meter)
	mil = konversiMeterMil(meter)
	fmt.Println(mil)
}
