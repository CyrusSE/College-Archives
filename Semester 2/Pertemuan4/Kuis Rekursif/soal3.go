package main

import "fmt"

func pangkat(a, b, start, hasil int) {
	if start == b {
		fmt.Println(hasil * a)
	} else {
		pangkat(a, b, start+1, hasil*a)
	}
}

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	pangkat(a, b, 1, 1)
}
