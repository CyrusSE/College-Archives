package main

import "fmt"

func warna(a, b string) string {
	if a == "merah" && b == "kuning" {
		return "orange"
	} else if a == "kuning" && b == "biru" {
		return "hijau"
	} else if a == "biru" && b == "merah" {
		return "ungu"
	} else if a == "kuning" && b == "merah" {
		return "orange"
	} else if a == "biru" && b == "kuning" {
		return "hijau"
	} else if a == "merah" && b == "biru" {
		return "ungu"
	}
	return ""
}

func main() {
	var w1, w2 string
	fmt.Scan(&w1)
	fmt.Scan(&w2)
	fmt.Println(warna(w1, w2))
}
