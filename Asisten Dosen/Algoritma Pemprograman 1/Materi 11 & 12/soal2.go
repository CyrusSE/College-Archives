package main

import "fmt"

func main() {
	var uang, dompet int
	for {
		fmt.Scan(&uang)
		if uang >= 0 {
			dompet = dompet + uang
		} else {
			dompet = dompet - (-uang)
		}
		if uang == 0 {
			break
		}
	}
	fmt.Println(dompet)
}