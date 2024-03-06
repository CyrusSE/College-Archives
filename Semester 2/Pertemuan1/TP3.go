package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)
	if x >= 0 && y >= 0 {
		fmt.Print("Kuadran I")
	} else if x < 0 && y >= 0 {
		fmt.Print("Kuadran II")
	} else if x >= 0 && y < 0 {
		fmt.Print("Kuadran IV")
	} else {
		fmt.Print("Kuadran III")
	}
}
