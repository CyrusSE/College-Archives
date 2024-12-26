package main

import "fmt"

func main() {
	var username, password string
	var hitung int

	fmt.Scan(&username, &password)
	for username != "admin" || password != "admin" {
		fmt.Scan(&username, &password)
		hitung = hitung + 1
	}
	fmt.Println(hitung)
	fmt.Println("Login berhasil")
}
