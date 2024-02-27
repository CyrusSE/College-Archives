package main

import "fmt"

func main() {
	var username, password string
	var hitung int
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	for username != "admin" || password != "admin" {
		fmt.Print("Username : ")
		fmt.Scan(&username)
		fmt.Print("Password : ")
		fmt.Scan(&password)
		hitung = hitung + 1
	}
	fmt.Println(hitung)
	fmt.Print("Login berhasil")
}
