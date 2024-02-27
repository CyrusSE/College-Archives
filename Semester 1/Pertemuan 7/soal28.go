package main

import "fmt"

func main() {
	var n int
	var duluan bool
	fmt.Print("Masukkan angka untuk N : ")
	fmt.Scan(&n)
	save := 1
	number := 0
	for i := 0; i <= n; i++ {
		duluan = (i%2 == 0)
		for duluan {
			fmt.Println(number)
			number = save + number
			break
		}
		for !duluan {
			fmt.Println(save)
			save = number + save
			break
		}
	}
}
