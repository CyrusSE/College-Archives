package main

import "fmt"

func manager(a, b, c, d, e string) string {
	if a == "kalah" && b == "kalah" && c == "kalah" && d == "kalah" && e == "kalah" {
		return "dipecat"
	} else {
		return "tidak dipecat"
	}
}

func main() {
	var p1, p2, p3, p4, p5 string
	fmt.Scan(&p1, &p2, &p3, &p4, &p5)
	fmt.Println(manager(p1, p2, p3, p4, p5))
}
