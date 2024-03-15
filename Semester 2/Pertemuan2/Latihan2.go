package main

import "fmt"

func validNIM(nim string) bool {
	if len(nim) != 8 {
		return false
	}
	if nim[:3] != "212" {
		return false
	}
	return true
}

func main() {
	var NIM string
	var valid bool
	fmt.Scan(&NIM)
	fmt.Println(len(NIM))
	valid = validNIM(NIM)
	fmt.Println(valid)
}
