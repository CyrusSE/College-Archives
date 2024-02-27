package main

import "fmt"

func main() {
	var temp int
	terbesar := 0
	terkecil := 0
	check := 0
	total := 0
	rata := 0.0
	for {
		fmt.Scan(&temp)
		if temp == 0 {
			check += 1
			if check == 2 {
				break
			}
		}
		if check == 1 {
			if temp != 0 {
				check = 0
			}
		}
		total++
		if temp > terbesar {
			terbesar = temp
		}
		if temp < terkecil {
			terkecil = temp
		}
		rata += float64(temp)
	}
	fmt.Println(terbesar)
	fmt.Println(terkecil)
	fmt.Println(rata / float64(total))
}
